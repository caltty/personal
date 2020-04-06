import React from 'react';
// eslint-disable-next-line 
import ReactDOM from 'react-dom';
import { Modal, Tooltip } from 'antd';
import { Button } from 'antd';
import { SettingOutlined } from '@ant-design/icons';
import './index.css';

function displaySuccess() {
    Modal.success({
        title: 'Success',
        content: 'Agron ldap sets the attribute success.',
    });
}

function displayError() {
    Modal.error({
        title: 'Error',
        content: 'Agron ladp fails to set the attribute.',
    });
}

function displayLogoutError() {
    Modal.error({
        title: 'Error',
        content: 'Agron ladp fails to logout.',
    });
}

export default class extends React.Component {
    dnName = React.createRef()
    attrType = React.createRef()
    attrName = React.createRef()
    handleClickLogout = async () => {
        const response = await fetch("https://localhost:9443/logout", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem("token")
            },
        });

        const result = await response.json();
        if (result["errorCode"] === '0') {
            localStorage.removeItem("token")
            window.location.href = "/"
        } else {
            localStorage.removeItem("token")
            displayLogoutError()
        }
    }

    handleClickApply = async () => {
        const data = {
            dnName: this.dnName.current.value,
            attrType: this.attrType.current.value,
            attrName: this.attrName.current.value,
        }
        const AttrName = [data.attrName];
        const message = {
            "dn": data.dnName,
            "attrType": data.attrType,
            "attrVals": AttrName
        }

        console.log(JSON.stringify(message));
        const response = await fetch("https://localhost:9443/modify-attribute", {
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

    handleClickSetting = async () => {
        window.location.href = "/serverConfig";
    }

    render() {
        return (
            <div>
                <div className="row">
                    <div className="pull-right"><button className="lougoutButton" onClick={this.handleClickLogout}>Logout</button></div>
                    <div className="pull-right"><Tooltip title= "Setting"><Button type="primary" shape="circle" icon={<SettingOutlined />} onClick={this.handleClickSetting} /></Tooltip></div>
                </div>
                <div className="row">
                    <label className="col-sm-3 control-label mdtext">Argon ldap attribute setting:</label>
                </div>
                <br />
                <div className="row">
                    <div className="form-horizontal">
                        <div className="form-group">
                            <label htmlFor="" className="col-sm-2 control-label">DN name:</label>
                            <div className="col-sm-4">
                                <input type="text"
                                    name="dnName"
                                    className="form-control"
                                    ref={this.dnName}
                                />
                            </div>
                        </div>
                        <div className="form-group">
                            <label htmlFor="" className="col-sm-2 control-label">Attribute type:</label>
                            <div className="col-sm-4">
                                <input type="text"
                                    name="attrType"
                                    className="form-control"
                                    ref={this.attrType}
                                />
                            </div>
                        </div>
                        <div className="form-group">
                            <label htmlFor="" className="col-sm-2 control-label">Attribute name:</label>
                            <div className="col-sm-4">
                                <input type="text"
                                    name="attrName"
                                    className="form-control"
                                    ref={this.attrName}
                                />
                            </div>
                        </div>
                    </div>
                </div>
                <div className="form-group">
                    <div className="col-sm-2 col-sm-offset-4">
                        <button className="btn btn-primary  btn-block outline:none;" onClick={this.handleClickApply} > Apply </button>
                    </div>
                </div>
            </div>
        )
    }
}