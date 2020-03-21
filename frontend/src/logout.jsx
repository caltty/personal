import React from 'react';
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
        if (result["errorCode"] == 0) {
            window.location.href = "/"
        }
        localStorage.removeItem("token")
    }
    render() {
        return (
            <div>
                <div className="row">
                    <button className="btn btn-primary pull-right" onClick={this.handleClickLogout}>Logout</button>
                </div>
                <div className="row">
                    <div className="text-center mdtext" >welcome to dsjkfldsjfkld</div>
                </div>
            </div>
        )
    }
}