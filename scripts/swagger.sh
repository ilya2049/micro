# Generate http rest api for the hasherapi service.
swagger generate server \
--target=hasherapi \
--spec=api/rest-api.yml \
--main-package=hasherapi \
--server-package=internal/system/restapi \
--model-package=internal/system/restapi/models \
--implementation-package=hasherapi/internal/system/restapi/handler \
--name="hasherapi"