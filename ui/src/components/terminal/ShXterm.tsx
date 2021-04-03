import React from 'react'
import { TerminalSingleton } from '../../lib/TerminalSingleton';

// https://gist.github.com/GitSquared/2049d7e85eaddeeeaa44e8404fe0b0e1

const styles = {
    terminal: {
        height: "97vh",
        padding: 3
    }
}

type XtermProperties = {
}
type XtermState = {
}

export class ShXterm extends React.Component<XtermProperties, XtermState> {
    private terminalRef: React.RefObject<HTMLDivElement>

    constructor(props: XtermProperties) {
        super(props)
        this.terminalRef = React.createRef()
    }

    async componentDidMount() {
        let term = TerminalSingleton.getInstance()
        await term.init()
       
        if (this.terminalRef.current) {
			// Creates the terminal within the container element.
            term.getTerminal().open(this.terminalRef.current)
            term.fit()
            term.getTerminal().focus()
        }
    }


    componentWillUnmount() {
		// When the component unmounts dispose of the terminal and all of its listeners.
		// this.terminal.dispose()
        // this.wconn?.closeConnection()
	}

    render() {
        return (
            <div style={styles.terminal as React.CSSProperties} ref={this.terminalRef}></div>
        )
    }
}