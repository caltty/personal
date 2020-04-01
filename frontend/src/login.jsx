import React from 'react';
// eslint-disable-next-line 
import ReactDOM, { render } from 'react-dom';
import './index.css';
import logo from './logo.svg';
import {Modal} from 'antd';

function display() {
    Modal.error({
        title: 'Error',
        content: 'Username or password is not correct',
      });
} 

export default class extends React.Component {
    username = React.createRef()
    password = React.createRef()
 
    handleClickLogin = async () => {
        const data = {
            username:this.username.current.value,
            password:this.password.current.value,
        }
        const message = {
            /**"username": "testuser1@hpdm.sh",
            "password": "Shanghai2010"*/
            "username": data.username,
            "password": data.password
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
            localStorage.setItem("username",this.username)
            console.log(localStorage.getItem("token"))
            window.location.href = "/logout"
        }else{
            display();
        }
    }
    render() {
        console.log(logo);
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
                                    <input type="text"
                                        name="username"
                                        className="form-control"
                                        ref={this.username}
                                         />
                                </div>
                            </div>

                            <div className="form-group">
                                <label htmlFor="password" className="col-sm-2 control-label">Password:</label>
                                <div className="col-sm-4">
                                    <input type="password"
                                        name="password"
                                        id="password"
                                        className="form-control"
                                        ref={this.password}
                                         />
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
