#!/bin/bash

CREDENTIAL_MANAGER_URL="127.0.0.1:4771"
IOT_URL="127.0.0.1:4772"


function getFieldFromJson(){
  local __json=$1
  local __field=$2

  local __fieldValue=$(jq -n "${__json}" | jq ".${__field}" | tr { '\n' | tr , '\n' | tr } '\n')


  if [[ -z "$__fieldValue" ]]; then
      echo "${__json}"
      exit 1
  fi

  __fieldValue="${__fieldValue//\"/""}"
  echo "${__fieldValue}"
}

function runIoTTestData(){
  local __credId1=$1
  local __credId2=$2
  local __credId3=$3

  echo "${__credId1}"
  echo "${__credId2}"
  echo "${__credId3}"


  curl -sb -X POST ${IOT_URL}/white-list -d "{\"door_id\": \"blue\",\"credential_id\":${__credId1}}"
  curl -sb -X POST ${IOT_URL}/white-list -d "{\"door_id\": \"blue\",\"credential_id\":${__credId2}}"
  curl -sb -X POST ${IOT_URL}/white-list -d "{\"door_id\": \"blue\",\"credential_id\":${__credId3}}"
  curl -sb -X POST ${IOT_URL}/white-list -d "{\"door_id\": \"green\",\"credential_id\":${__credId2}}"
  curl -sb -X POST ${IOT_URL}/white-list -d "{\"door_id\": \"red\",\"credential_id\":${__credId3}}"

  echo ""
  echo "WhiteList: "
  curl -sb -X GET ${IOT_URL}/white-list

  curl -sb -X DELETE ${IOT_URL}/white-list -d "{\"door_id\": \"red\",\"credential_id\":${__credId3}}"

  echo ""
  echo ""
  echo "WhiteList after delete: "
  curl -sb -X GET ${IOT_URL}/white-list
}

function runCredentialManagerTestData(){
  local __responseCreateUser=$(curl -sb -X POST ${CREDENTIAL_MANAGER_URL}/user -d '{"name": "Rafal Bil"}')

  local __userId=$(getFieldFromJson "${__responseCreateUser}" "Id")
  echo "UserId: ${__userId}"

  local __credential1Id=$(curl -sb -X POST ${CREDENTIAL_MANAGER_URL}/credential -d '{"credential": "12345678"}')

  echo "Credential1: ${__credential1Id}"

  local __credential2Id=$(curl -sb -X POST ${CREDENTIAL_MANAGER_URL}/credential -d '{"credential": "98765432"}')
  echo "Credential2: ${__credential2Id}"

  local __credential3Id=$(curl -sb -X POST ${CREDENTIAL_MANAGER_URL}/credential -d '{"credential": "91827364"}')
  echo "Credential3: ${__credential3Id}"

  local __userCredential1=$(curl -sb -X POST ${CREDENTIAL_MANAGER_URL}/user-credential -d "{\"user_id\":\"${__userId}\",\"credential_id\":${__credential1Id}}")

  local __userCredential2=$(curl -sb -X POST ${CREDENTIAL_MANAGER_URL}/user-credential -d "{\"user_id\":\"${__userId}\",\"credential_id\":${__credential2Id}}")

  local __userCredential3=$(curl -sb -X POST ${CREDENTIAL_MANAGER_URL}/user-credential -d "{\"user_id\":\"${__userId}\",\"credential_id\":${__credential3Id}}")

  runIoTTestData ${__credential1Id} ${__credential2Id} ${__credential3Id}

  curl -v -X POST 127.0.0.1:4772/access -d '{"door_id":"blue","credential":"12345678"}'
}


runCredentialManagerTestData
