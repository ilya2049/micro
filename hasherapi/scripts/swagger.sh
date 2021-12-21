swagger generate server \
--spec=api/rest-api.yml \
--main-package=hasherapi \
--server-package=internal/system/restapi \
--model-package=internal/system/restapi/models \
--implementation-package=hasherapi/internal/system/restapi/handler \
--name="hasherapi"