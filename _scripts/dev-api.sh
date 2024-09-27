#!/bin/bash

SVC_CREDENTIAL_MANAGER_URL="127.0.0.1:4771"
SVC_IOT_URL="127.0.0.1:4772"



# Commands
apiCredentialManagerCommands(){

  local _apiCommand="$1"

  case "${_apiCommand}" in
    create-user)
      local _userName=$2
      local _jsonPayload="{\"name\":\"${_userName}\"}"
      curl -X POST -v -d "${_jsonPayload}" ${SVC_CREDENTIAL_MANAGER_URL}/user
      ;;
    create-credential)
      local _credentialCode=$2
      local _jsonPayload="{\"credential\":\"${_credentialCode}\"}"
      curl -X POST -v -d "${_jsonPayload}" ${SVC_CREDENTIAL_MANAGER_URL}/credential
      ;;
    assign-credential)
      local _userId=$2
      local _credentialId=$3
      local _jsonPayload="{\"user_id\":\"${_userId}\",\"credential_id\":\"${_credentialId}\"}"
      curl -X POST -v -d "${_jsonPayload}" ${SVC_CREDENTIAL_MANAGER_URL}/user-credential
    ;;
    authorize)
      local _doorId=$2
      local _credentialId=$3

      curl -X POST -v -d "${_jsonPayload}" ${SVC_CREDENTIAL_MANAGER_URL}/authorize/"${_doorId}"/"${_credentialId}"
    ;;
    revoke)
      local _doorId=$2
      local _credentialId=$3
      # shellcheck disable=SC2086
      curl -X DELETE -v -d "${_jsonPayload}" ${SVC_CREDENTIAL_MANAGER_URL}/authorize/${_doorId}/${_credentialId}
    ;;
    samples)
      cat <<-EOF
        dev-api.sh credential-manager create-user 'John Doe'
        dev-api.sh credential-manager create-credential 12345678
        dev-api.sh credential-manager assign-credential user-Id-xxx credential-Id-yyy
        dev-api.sh credential-manager authorize blue-door credential-Id-yyy
        dev-api.sh credential-manager revoke blue-door credential-Id-yyy

EOF
      ;;
    *)
      echo "Not supported command => [${_apiCommand}]"
      echo "Samples ==>>"
      dev-api.sh clients samples
      exit 255
  esac
}

apiIoTCommands(){

  local _apiCommand="$1"

  case "${_apiCommand}" in
    create-whitelist)
      local _doorId=$2
      local _credentialId=$3
      local _jsonPayload="{\"door_id\":\"${_doorId}\",\"credential_id\":\"${_credentialId}\"}"
      curl -X POST -v -d "${_jsonPayload}" ${SVC_IOT_URL}/white-listuser
      ;;
    delete-whitelist)
      local _doorId=$2
      local _credentialId=$3
      local _jsonPayload="{\"door_id\":\"${_doorId}\",\"credential_id\":\"${_credentialId}\"}"
      curl -X DELETE -v -d "${_jsonPayload}" ${SVC_IOT_URL}/white-listuser
      ;;
    request-access)
      local _doorId=$2
      local _credentialCode=$3
      local _jsonPayload="{\"door_id\":\"${_doorId}\",\"credential\":\"${_credentialCode}\"}"
      echo "${_jsonPayload}"
      curl -X POST -v -d "${_jsonPayload}" ${SVC_IOT_URL}/access
      ;;

    samples)
      cat <<-EOF
        dev-api.sh iot create-whitelist blue-door credential-Id-yyy
        dev-api.sh iot delete-whitelist blue-door credential-Id-yyy
        dev-api.sh iot request-access blue-door 12345678
EOF
      ;;
    *)
      echo "Not supported command => [${_apiCommand}]"
      echo "Samples ==>>"
      dev-api.sh clients samples
      exit 255
  esac
}


#
# Main program
#
API_BLOCK="$1"
shift 1

case "${API_BLOCK}" in
  credential-manager)
    apiCredentialManagerCommands "$@"
    ;;
  iot)
    apiIoTCommands "$@"
    ;;
  *)
      cat <<-EOF
  Not supported Api => [${API_BLOCK}]
  Supported commands:
    CREDENTIAL-MANAGER-SVC =>
      - dev-api.sh credential-manager create-user
      - dev-api.sh credential-manager create-credential
      - dev-api.sh credential-manager assign-credential
      - dev-api.sh credential-manager authorize
      - dev-api.sh credential-manager revoke

    IOT-SVC =>
      - dev-api.sh iot create-whitelist
      - dev-api.sh iot delete-whitelist
      - dev-api.sh iot request-access
EOF
    exit 255
esac