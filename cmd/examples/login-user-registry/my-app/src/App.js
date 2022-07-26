import React, { useState, useEffect } from "react";
import Modal from 'react-modal';
import { ethers } from 'ethers';
import './App.css';
import google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb.js'
import { UserList , User , PubKey , UserName , LoginResp } from './identity/login_pb'
import { UserRegistryLoginPromiseClient } from './identity/login_grpc_web_pb'

var client = new UserRegistryLoginPromiseClient('http://localhost:8080')

const customStyles = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-50%',
    transform: 'translate(-50%, -50%)',
    textalign:'center',

  },
};

function App() {
  const [name, setName] = useState("")
  const [privKeyHex, setPrivKeyHex] = useState("")
  const [message, setMessage] = useState("")
  const [modalIsOpen, setIsOpen] = React.useState(false);

  useEffect(() => {
    client.listAllUsers(empty, {}, (err, response) => {
      console.log(response.getUsersList())
      if (err) {
        console.log(err)
      }
      
    })
  });

  function openModal() {
    setIsOpen(true);
  }

  function closeModal() {
    setIsOpen(false);
  }


  const getName = (event) => {
    setName(event.target.value)
  };

  const getPrivKey = (event) => {
    setPrivKeyHex(event.target.value)
  };

  const getMessage = (message) => {
    setMessage(message)
  }

  const empty = new google_protobuf_empty_pb.Empty;
  const verifyLogin = async (event) => {

    const privKey = new ethers.utils.SigningKey(privKeyHex);
    const userName = new UserName();
    userName.setUsername(name);
    
    var loginInfo = await client.requestLogin(userName, {})

    const signature = privKey.signDigest(ethers.utils.id(loginInfo.getUnsignedmsg()))
    const jointSig = ethers.utils.joinSignature(signature)
    const sig = ethers.utils.arrayify(jointSig)
    loginInfo.setSignature(sig)  

    var loginResp = await client.login(loginInfo, {})
    var status = loginResp.getStatus()
    if (status === "ok") {
      getMessage("Login Success!")
    } else {
      getMessage("Login Failed")
    }

    openModal()
  }


  return ( 
    <div className="App">
      <h1>Welcome</h1>
      <div className="User-Name-Wrapper">
        <input type="text" className="User-Name" placeholder="Username" onChange={getName}/>
      </div>
      <div className="Priv-Key-Wrapper">
        <input type="text" className="Priv-Key" placeholder="Private Key" onChange={getPrivKey}/>
      </div>
      <div className="Login-Button-Wrapper">
        <button
            type="submit"
            value="Submit"
            onClick={verifyLogin}
            className="Login-Button">
        Login
        </button>
        
      <Modal
        isOpen={modalIsOpen}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel="Example Modal"
        ariaHideApp={false}
      >
        <div className="Message">{message}</div>
      <div className="Close-Wrapper">
        <button className="Close" onClick={closeModal}>close</button>
      </div>
      </Modal>
   

      </div>
    </div>
  );
} 

export default App;

    /*
    var loginInfo = await client.requestLogin(userName, {}, async (err, response) => {
      await response
      console.log("debug0", response)
      if (err) {
        console.log(err)
        return null
      }
      return response
    })
 
    console.log("debug1", await(loginInfo))
  */
    /*
    const requestLogin = async () => {
   
      client.requestLogin(userName, {}, (err, response) => {  
        console.log("debug0", response)
        if (err) {
          console.log(err)
        }
        loginInfo = response
      })
      
      loginInfo = await client.requestLogin(userName, {})
    }
    await requestLogin()
    */

     
     /*
     return await client.requestLogin(userName, {})
    }
   console.log("debug1", requestLogin())
  
   
    const login = async () => {
      loginInfo = await requestLogin()
      console.log("debug2", loginInfo)
      const signature = ethers.utils.joinSignature(privKey.signDigest(ethers.utils.id(loginInfo.getUnsignedmsg())))
      loginInfo.setSignature(signature)

      var loginResp = new LoginResp()
      client.login(loginInfo, {}, (err, response) => {
        loginResp = response
        if (err) {
          console.log(err)
        }
      })
    }

    login()
   
    /*
    requestLogin().then( 
      () => {
      console.log("debug2", loginInfo)
      const signature = ethers.utils.joinSignature(privKey.signDigest(ethers.utils.id(loginInfo.getUnsignedmsg())))
      loginInfo.setSignature(signature)
  
      var loginResp = new LoginResp()
      client.login(loginInfo, {}, (err, response) => {
        loginResp = response
        if (err) {
          console.log(err)
        }
      })
      } 
    )
    */

        /*
    const sign = crypto.createSign("DSA-SHA512")
    sign.update(loginInfo.getUnsignedmsg())
    //const signature = sign.sign(privKeyHex, 'base64')
    loginInfo.setSignature(signature)
    console.log("signature", signature)
    */
    /*
    let dataHash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes(JSON.stringify(loginInfo.getUnsignedmsg())))
    let dataHashBin = ethers.utils.arrayify(dataHash)
    let wallet = new ethers.Wallet(privKeyHex)
    let signature = await wallet.signMessage(dataHashBin)
    loginInfo.setSignature(signature)
    console.log(signature)
    */
    