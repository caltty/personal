import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import logo from './logo.svg';

export default class extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: '',
            errorCode: '',
            token: ''
        }
    }

    handleChangeUserName(e) {
        this.setState({
            username: e.target.value
        })
    }

    handleChangePassword(e) {
        this.setState({
            password: e.target.value
        })
    }

    async handleClickLogin() {

        const message = {
            "username": "CN=testuser2,OU=argonldap,OU=shanghai,DC=sh,DC=argon",
            "password": "Shanghai2010"
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

        if (errorCode == 0) {
            window.location.href = "/logout"
        }

        localStorage.setItem("token", 'Bearer ' + result["token"])
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
                                        onChange={this.handleChangeUserName.bind(this)}
                                        defaultValue={this.state.username} />
                                </div>
                            </div>

                            <div className="form-group">
                                <label htmlFor="password" className="col-sm-2 control-label">Password:</label>
                                <div className="col-sm-4">
                                    <input type="password"
                                        name="password"
                                        id="password"
                                        className="form-control"
                                        onChange={this.handleChangePassword.bind(this)}
                                        defaultValue={this.state.password} />
                                </div>
                            </div>
                            <div className="form-group">
                                <div className="col-sm-2 col-sm-offset-4">
                                    <button className="btn btn-primary  btn-block" onClick={this.handleClickLogin} > login </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}