import React from 'react';
// eslint-disable-next-line 
import ReactDOM from 'react-dom';
import { Modal } from 'antd';
import { Input } from 'antd';
import Config from "./config"
import { getServerConfig } from "./api"

function displaySuccess() {
    Modal.success({
        title: 'Success',
        content: 'Agron ldap server configurations are setted successfully.',
    });
}

function displayError() {
    Modal.error({
        title: 'Error',
        content: 'Agron ldap server configurations are fail to setted.',
    });
}


// eslint-disable-next-line
export default class ServerConfig extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            server: {},
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

        getServerConfig().then(p => {
            this.setState({ server: p.data })
        })
    }

    onChange = e => {
        let { server } = this.state
        server[e.target.name] = e.target.value
        this.setState({ server })
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

    handleSave = async () => {
        const message = {
            "server": {
                "port": this.state.server.port,
                "sslCertificateFile": this.state.server.sslCertifcateFile,
                "sslKeyFile": this.state.server.sslKeyFile
            },
            "ldap": {
                "url": this.state.ldap.url,
                "bindUser": this.state.ldap.bindUser,
                "bindPassword": this.state.ldap.bindPassword,
                "baseDn": this.state.ldap.baseDn
            },
            "jwtAuthentication": {
                "tokenDuration": this.state.jwtAuthentcation.tokenDuration,
                "secretKeyManagerType": this.state.jwtAuthentcation.secretKeyManagerType,
                "memorySecretKeyManagerSettings": {
                    "cleanupInterval": this.state.jwtAuthentcation.memorySecretKeyManagerSettings.cleanUpInterval
                },
                "redisSecretKeyManagerSettings": {
                    "serverAddress": this.state.redisSecretKeyManagerSettings.serverAddress,
                    "maxIdle": this.state.redisSecretKeyManagerSettings.maxIdle,
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
        const { server } = this.state;
        //this.getConfigurations();
        return (
            <Config>
                <h1>Agron Ldap Server</h1>
                <div className="form-horizontal">
                    <div className="form-group">
                        <label htmlFor="" className="control-label col-sm-2">Port:</label>
                        <div className="col-sm-4">
                            <Input placeholder="port" name="port" value={server.port} onChange={this.onChange} />
                        </div>
                    </div>
                    <div className="form-group">
                        <label htmlFor="" className="control-label col-sm-2">Ssl certifcateFile:</label>
                        <div className="col-sm-4">
                            <Input placeholder="sslCertifcateFile" name="certFile" value={server.certFile} onChange={this.onChange} />
                        </div>
                    </div>
                    <div className="form-group">
                        <label htmlFor="" className="control-label col-sm-2">Ssl key file:</label>
                        <div className="col-sm-4">
                            <Input placeholder="sslKeyFile" name="sslKey" value={server.sslKey} onChange={this.onChange} />
                        </div>
                    </div>
                    <div className="form-group">
                        <div className="col-sm-2 col-sm-offset-4">
                            <button className="btn btn-primary  btn-block outline:none;" onClick={this.handleSave} > Save </button>
                        </div>
                    </div>
                </div>
            </Config>
        )
    }
}