import React from 'react'
import { TerminalView } from './view/TerminalView'
import { Layout, Menu } from 'antd';
import { ExperimentOutlined, CodeOutlined } from '@ant-design/icons';
import {
    Switch,
    Route,
    withRouter,
    RouteComponentProps
  } from "react-router-dom";

const styles = {
    header: {
        padding: 0,
        width: "100%",
        zIndex: 100,
        height: 40,
        lineHeight: 3,
        // paddingLeft: 220,
        position: 'fixed',
        overflow: "hidden",
        borderBottom: "1px solid #666",

    } as React.CSSProperties,
    
    mainContent: {
        // marginTop: 40,
        marginTop: 0,
        marginLeft: 0, 
        marginRight: 0
        // overflow: 'initial' 
    } as React.CSSProperties
}

type ContainerState = {
    menuCurrent: string
}

type ContainerProps = {
}

interface IComponentProps extends RouteComponentProps {
}

type PropsType = IComponentProps & ContainerProps

class MainContainer extends React.Component<PropsType, ContainerState> {

    private viewNames = [
        "Shell"
    ]
    private icons = [
        <ExperimentOutlined />,
        <CodeOutlined />
    ]
    private urls = [
        "/",
    ]

    state: ContainerState = {
        menuCurrent: "0"
    }

    componentDidMount() {
        const pathname = this.props.location.pathname
        for (let idx in this.urls) {
            if (pathname === this.urls[idx]) {
                this.setState({
                    menuCurrent: idx.toString()
                })
                break
            }
        }
    }

    private handleMenuClick = (e: any) => {
        const { history } = this.props
        this.setState({ 
            menuCurrent: e.key 
        })
        const url = this.urls[parseInt(e.key)]
        history.push(url)
    }
    
    render() {
        return (
            <React.Fragment>
                <Layout>
                    {/* <Layout.Header style={styles.header}>
                        <Menu theme="dark" 
                            selectedKeys={[this.state.menuCurrent]}
                            onClick={this.handleMenuClick}
                            mode="horizontal" 
                            defaultSelectedKeys={["0"]}>
                            {this.viewNames.map( (item, idx) => {
                                return (
                                    <Menu.Item key={idx} icon={this.icons[idx]}>{item}</Menu.Item>
                                )
                            })}
                        </Menu>
                    </Layout.Header> */}
                    <Layout.Content style={styles.mainContent}>
                        <Switch>
                            <Route exact path="/">
                                <TerminalView />
                            </Route>
                        </Switch>
                    </Layout.Content>
                </Layout>
            </React.Fragment>
        )
    }
}

export const MainContainerWithRouter = withRouter(MainContainer)