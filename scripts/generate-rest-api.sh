# Generate http rest api for the hasherapi service.
swagger generate server \
--target=hasherapi \
--spec=api/rest-api.yml \
--main-package=hasherapi \
--server-package=system/restapi \
--model-package=system/restapi/models \
--implementation-package=hasherapi/system/restapi/handler \
--name="hasherapi"