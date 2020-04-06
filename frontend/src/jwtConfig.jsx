import React from 'react';
// eslint-disable-next-line 
import ReactDOM from 'react-dom';
import { Modal } from 'antd';
import {Input } from 'antd';
import Config from "./config"

function displaySuccess() {
    Modal.success({
        title: 'Success',
        content: 'Agron ldap JWT configurations are setted successfully.',
    });
}

function displayError() {
    Modal.error({
        title: 'Error',
        content: 'Agron ldap JWT configurations are fail to setted.',
    });
}


// eslint-disable-next-line
export default class LdapConfig extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            server: {
                port: "9443",
                sslCertifcateFile: "server.cert",
                sslKeyFile: "server.key"
            },

            ldap: {
                url: "ldap://15.15.182.177:389",
                bindUser: "Administrator@hpdm.sh",
                bindPassword: "Shanghai2010",
                baseDn: "dc=hpdm,dc=sh"
            },

            jwtAuthentcation: {
                tokenDuration: "30",
                secretKeyManagerType: "memory",
                memorySecretKeyManagerSettings: {
                    cleanUpInterval: "1"
                }
            },

            redisSecretKeyManagerSettings: {
                serverAddress: "192.168.153.239:6389",
                maxIdle: "10",
                MaxActive: "50",
                IdleTimeout: "10"
            }
        }
    }

    onChange = e =>{
        this.setState({[e.target.name]:e.target.value})
    }

    getConfigurations = async () => {
        const response = await fetch("https://localhost:9443/config", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem("token")
            },
        });

        const result = await response.json();
        this.state.server.port = result["config"]["server"]["port"];
        this.state.server.sslCertifcateFile = result["config"]["server"]["sslCertifcateFile"];
        this.state.server.sslKeyFile = result["config"]["server"]["sslKeyFile"];
        this.state.ldap.url = result["config"]["ldap"]["url"];
        this.state.ldap.bindUser = result["config"]["ldap"]["bindUser"];
        this.state.ldap.bindPassword = result["config"]["ldap"]["bindPassword"];
        this.state.ldap.baseDn = result["config"]["ldap"]["baseDn"];
        this.state.jwtAuthentcation.secretKeyManagerType = result["config"]["jwtAuthentcation"]["secretKeyManagerType"];
        this.state.jwtAuthentcation.tokenDuration = result["config"]["jwtAuthentcation"]["tokenDuration"];
        this.state.jwtAuthentcation.memorySecretKeyManagerSettings.cleanUpInterval = result["config"]["jwtAuthentcation"]["memorySecretKeyManagerSettings"]["cleanUpInterval"];
        this.state.redisSecretKeyManagerSettings.serverAddress = result["config"]["redisSecretKeyManagerSettings"]["serverAddress"];
        this.state.redisSecretKeyManagerSettings.maxIdle = result["config"]["redisSecretKeyManagerSettings"]["maxIdle"];
        this.state.redisSecretKeyManagerSettings.MaxActive = result["config"]["redisSecretKeyManagerSettings"]["MaxActive"];
        this.state.redisSecretKeyManagerSettings.IdleTimeout = result["config"]["redisSecretKeyManagerSettings"]["IdleTimeout"];

    }

    handleSave = async() =>{
        const message = {
            "server": {
                "port": this.state.server.port,
                "sslCertificateFile": this.state.server.sslCertifcateFile,
                "sslKeyFile": this.state.server.sslKeyFile
            },
            "ldap" : {
                "url" : this.state.ldap.url,
                "bindUser" : this.state.ldap.bindUser,
                "bindPassword" : this.state.ldap.bindPassword,
                "baseDn" : this.state.ldap.baseDn
            },
            "jwtAuthentication" : {
                "tokenDuration" : this.state.jwtAuthentcation.tokenDuration,
                "secretKeyManagerType" : this.state.jwtAuthentcation.secretKeyManagerType,
                "memorySecretKeyManagerSettings" : { 
                    "cleanupInterval" : this.state.jwtAuthentcation.memorySecretKeyManagerSettings.cleanUpInterval
                },
                "redisSecretKeyManagerSettings" : { 
                    "serverAddress": this.state.redisSecretKeyManagerSettings.serverAddress,
                    "maxIdle" : this.state.redisSecretKeyManagerSettings.maxIdle,
                    "MaxActive": this.state.redisSecretKeyManagerSettings.MaxActive,
                    "IdleTimeout": this.state.redisSecretKeyManagerSettings.IdleTimeout
                }
            }
        }
        const response = await fetch("https://localhost:9443/config", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem("token")
            },
            body: JSON.stringify(message)
        });

        const result = await response.json();
        const errorCode = result;
        if (errorCode === '0') {
            displaySuccess()
        } else {
            displayError()
        }
    }

    render() {
        const {tokenDuration,secretKeyManagerType,cleanUpInterval} = this.state;
        //this.getConfigurations();
        return (
            <Config>
                <h1>Agron Ldap JWT Configuration</h1>
                <br/>
                <div className="form-horizontal">
                    <div className="form-group">
                        <label htmlFor="" className="control-label col-sm-3">Token duration:</label>
                        <div className="col-sm-4">
                            <Input placeholder="token duration" name="token duration"   value={tokenDuration} onChange={this.onChange} />
                        </div>
                    </div>
                    <div className="form-group">
                        <label htmlFor="" className="control-label col-sm-3">SecretKey manager type:</label>
                        <div className="col-sm-4">
                            <Input placeholder="secretKey manager type" name="secretKey manager type"   value={secretKeyManagerType} onChange={this.onChange} />
                        </div>
                    </div>
                    <label>Memory secret key manager settings:</label>
                    <div className="form-group">
                        <label htmlFor="" className="control-label col-sm-3">Cleanup interval:</label>
                        <div className="col-sm-4">
                            <Input placeholder="cleanup interval" name="cleanup interval"   value={cleanUpInterval} onChange={this.onChange} />
                        </div>
                    </div>
                    <div className="form-group">
                                <div className="col-sm-2 col-sm-offset-5">
                                    <button className="btn btn-primary  btn-block outline:none;" onClick={this.handleSave} > Save </button>
                                </div>
                            </div>
                </div>
            </Config>
        )
    }
}