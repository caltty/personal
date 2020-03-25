import React from 'react';
// eslint-disable-next-line 
import ReactDOM from 'react-dom';

export default class extends React.Component {
    async handleClickLogout() {
        const response = await fetch("https://localhost:9443/logout", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem("token")
            },
        });

        const result = await response.json();
        if (result["errorCode"] === '0') {
            window.location.href = "/"
        }
        localStorage.removeItem("token")
    }
    render() {
        return (
            <div>
                <div className="row">
                    <div className="pull-right col-md-1"><button className="btn btn-primary outline:none;" onClick={this.handleClickLogout}>Logout</button></div>
                </div>
                <div className="row">
                    <div className="text-center mdtext" >Welcome to argon ladp authentication</div>
                </div>
            </div>
        )
    }
}