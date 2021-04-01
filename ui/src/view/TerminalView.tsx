import React from 'react'

import { Layout, Row, Col } from 'antd';
import { ShXterm } from '../components/terminal/ShXterm';


const styles = {
    sider: {
        overflow: 'auto',
        height: '100vh',
        position: 'fixed',
        left: 0,
    } as React.CSSProperties,

    mainContent: {
    } as React.CSSProperties,
    
    terminalContainer: {
        marginTop: 10, 
        padding: 0,
        borderStyle: "solid",
        borderWidth: "1px",
        borderColor: "#666",
        backgroundColor: "#222",
    } as React.CSSProperties
}

type TerminalViewState = {
    connected: boolean
}

export class TerminalView extends React.Component<{}, TerminalViewState> {
    state: TerminalViewState = {
        connected: false
    }

    async componentDidMount() {
        
    }


    componentWillUnmount() {
    }

    render() {
        return (
            <Layout>
                <Layout style={styles.mainContent}>
                    <Layout.Content style={{ margin: '0px 5px 0', overflow: 'initial' }}>
                    <Row>
                        <Col span={24}>
                            <div style={styles.terminalContainer}>
                                <ShXterm />
                            </div>
                        </Col>
                    </Row>
                    </Layout.Content>
                </Layout>
            </Layout>
            
        )
    }
}