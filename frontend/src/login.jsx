import React from 'react';
// eslint-disable-next-line 
import ReactDOM, { render } from 'react-dom';
import './index.css';
import logo from './logo.svg';
import { Modal, Input  } from 'antd';
import { UserOutlined } from '@ant-design/icons';

function display() {
    Modal.error({
        title: 'Error',
        content: 'Username or password is not correct',
    });
}

export default class extends React.Component {
    constructor(props){
        super(props);
        this.state = {
            username:"",
            password:"",
        }
    }
    
    onChange = e =>{
        this.setState({[e.target.name]:e.target.value})
    }

    handleClickLogin = async () => {
        const message = {
            /**"username": "testuser1@hpdm.sh",
            "password": "Shanghai2010"*/
            "username": this.state.username,
            "password": this.state.password
        }
        const response = await fetch("https://localhost:9443/login", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem("token")
            },
            body: JSON.stringify(message)
        });

        const result = await response.json();
        const { token, errorCode } = result;
        console.log(errorCode);
        if (errorCode === '0') {
            localStorage.setItem("token", 'Bearer ' + token)
            console.log(localStorage.getItem("token"))
            window.location.href = "/logout"
        } else {
            display();
        }
    }
    render() {
        const {username,password} = this.state;
        return (
            <div className="container">
                <br />
                <br />
                <br />
                <div className="row">
                    <div className="col-sm-3 col-sm-offset-4">
                        <img src={logo} className="Auth-logo img-responsive" alt="logo" />
                    </div>

                </div>
                <br />
                <div className="row">
                    <div className="col-sm-offset-3">
                        <div className="form-horizontal">
                            <div className="form-group">
                                <label htmlFor="" className="control-label col-sm-2">User Name:</label>
                                <div className="col-sm-4">
                                    <Input placeholder="username" name="username" prefix={<UserOutlined />} className="autocomplete:false" value = {username} onChange ={this.onChange}/>
                                </div>
                            </div>

                            <div className="form-group">
                                <label htmlFor="password" className="col-sm-2 control-label">Password:</label>
                                <div className="col-sm-4">
                                    <Input.Password name="password" placeholder="input password" className="autocomplete:false" value = {password} onChange ={this.onChange}/>
                                </div>
                            </div>
                            <div className="form-group">
                                <div className="col-sm-2 col-sm-offset-4">
                                    <button className="btn btn-primary  btn-block outline:none;" onClick={this.handleClickLogin} > Login </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
