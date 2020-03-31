default: netapp.yaml
	GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w" openapi-generator generate -i netapp.yaml -g go -o ontap --package-name ontap

remove-dupe-opts: ## Removes duplicate Opt structs from the generated code
	@for struct in $$(grep -h 'type .\{1,\} struct' ontap/*.go | grep Opts  | sort | uniq -c | grep -v '^      1' | awk '{print $$3}'); do \
	  for f in $$(/bin/ls ontap/*.go); do \
	    if grep -qF "type $${struct} struct" "$${f}"; then \
	      if eval "test -z \$${$${struct}}"; then \
	        echo "skipping first appearance of $${struct} in file $${f}"; \
	        eval "export $${struct}=1"; \
	      else \
	        echo "removing dupe $${struct} from file $${f}"; \
	        tr '\n' '\r' <"$${f}" | sed 's~// '"$${struct}"'.\{1,\}type '"$${struct}"' struct {[^}]\{1,\}}~~' | tr '\r' '\n' >"$${f}.tmp"; \
	        mv -f "$${f}.tmp" "$${f}"; \
	      fi; \
	    fi \
	  done \
	done
	gofmt -w .
