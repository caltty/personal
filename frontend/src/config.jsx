import React from 'react';
// eslint-disable-next-line 
import ReactDOM from 'react-dom';
import { Tooltip } from 'antd';
import { Button } from 'antd';
import { ArrowRightOutlined, ReadOutlined, ReconciliationOutlined, ProjectOutlined, ProfileOutlined } from '@ant-design/icons';
import { Menu, Layout } from 'antd';
import './config.css';
import {Link} from 'react-router-dom'

const { Header, Content, Sider } = Layout;
export default class extends React.Component {


    handleClick = e => {
        console.log('click ', e);
    };



    handleClickBack = async () => {
        window.location.href = "/logout";
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
        } else {
        }
    }

    render() {
        //this.getConfigurations();
        return (
            <Layout style={{ minHeight: '100%' }}>
                <Header className="header">
                    <div className="logo" />
                    <h2 style={{ color: '#fff' }}>Agron Ldap Configuration
                        <div className="pull-right">
                            <div><Tooltip title="Back"><Button type="primary" shape="circle" icon={<ArrowRightOutlined />} onClick={this.handleClickBack} /></Tooltip></div>
                        </div>
                    </h2>
                </Header>
                <Layout>
                    <Sider width={330} className="site-layout-background" >
                        <Menu
                            mode="inline"
                            defaultOpenKeys={['1']}
                            style={{ height: '100%', borderRight: 0 }}
                        >
                            <Menu.Item key="1"><Link to="/serverConfig"><ProjectOutlined />Server</Link></Menu.Item>
                            <Menu.Item key="2"><Link to="/ldapConfig"><ReadOutlined />Ldap</Link></Menu.Item>
                            <Menu.Item key="3"><Link to="/jwtConfig"><ReconciliationOutlined />jwt Authorization</Link></Menu.Item>
                            <Menu.Item key="4"><Link to="/redisConfig"><ProfileOutlined />Redis secret key manager settings</Link></Menu.Item>
                        </Menu>
                    </Sider>
                    <Layout style={{ padding: '24px' }}>
                        <Content
                            className="site-layout-background"
                            style={{
                                padding: 24,
                                margin: 0,
                                minHeight: 280,
                            }}
                        >
                            {this.props.children}
                        </Content>
                    </Layout>
                </Layout>
            </Layout>
        )
    }
}