



peer lifecycle chaincode install token-org1.tgz

Chaincode code package identifier: token:2a69a1a6bd64cce249d464983de13bf8f47e194f940b645ff20a608aab37eda5


peer lifecycle chaincode install token-org2.tgz

Chaincode code package identifier: token:447db05fde5db734b60018782c2a8f987efb5c277d73aa26fed0d52a79b4203b

--approve---

peer lifecycle chaincode approveformyorg \
  --channelID thoaitest \
  --name token \
  --version 1.0 \
  --init-required \
  --package-id token:2a69a1a6bd64cce249d464983de13bf8f47e194f940b645ff20a608aab37eda5 \
  --sequence 1 \
  -o orderer:7050 \
  --tls \
  --cafile $ORDERER_CA


peer lifecycle chaincode approveformyorg \
  --channelID thoaitest \
  --name token \
  --version 1.0 \
  --init-required \
  --package-id token:447db05fde5db734b60018782c2a8f987efb5c277d73aa26fed0d52a79b4203b \
  --sequence 1 \
  -o orderer:7050 \
  --tls \
  --cafile $ORDERER_CA


  peer lifecycle chaincode checkcommitreadiness \
    --channelID thoaitest \
    --name token \
    --version 1.0 \
    --sequence 1 \
    --init-required \
    --output json



    -- Commit chaincode ---
peer lifecycle chaincode commit \
  --channelID thoaitest \
  --name token \
  --version 1.0 \
  --sequence 1 \
  --init-required \
  --tls \
  --peerAddresses peer0-org1:7051 \
    --tlsRootCertFiles /organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
  --peerAddresses peer0-org2:7051 \
    --tlsRootCertFiles /organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
  --tls \
  --cafile $ORDERER_CA


  peer chaincode invoke \
  -o orderer:7050 \
  --tls \
  --cafile $ORDERER_CA \
  -C thoaitest \
  -n token \
  --isInit \
  -c '{"Args":["Initialize"]}' \
  --peerAddresses peer0-org1:7051 \
  --tlsRootCertFiles /organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt