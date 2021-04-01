import { Terminal } from 'xterm';
import 'xterm/css/xterm.css'
import { FitAddon } from 'xterm-addon-fit';
import { AttachAddon } from 'xterm-addon-attach';
import { getWsURL, sleep } from './Utils';

export class TerminalSingleton {
    private socket: WebSocket | null
    private pingInterval!: number
    private connected: boolean = false
    
    
    private static instance: TerminalSingleton | null = null
    
    private terminal!: Terminal 
    private fitAddon!: FitAddon
    private attachAddon!: AttachAddon
    
    public static getInstance(): TerminalSingleton {
        if (TerminalSingleton.instance !== null) {
            return TerminalSingleton.instance
        }
        TerminalSingleton.instance = new TerminalSingleton()
        return TerminalSingleton.instance
        
    }
    
    public getTerminal(): Terminal {
        return this.terminal
    }
    
    public fit() {
        this.fitAddon.fit()
    }
    
    private constructor() {
        this.socket = null
        this.terminal = new Terminal({
            convertEol: true,
            fontFamily: `'Fira Mono', monospace`,
            fontSize: 16,
        })
        
        this.terminal.setOption('theme', {
            background: "#222",
            foreground: "#fff",
        })
        
        this.fitAddon = new FitAddon()
        this.terminal.loadAddon(this.fitAddon)
        
    }
    
    public isConnected(): boolean {
        return this.connected
    }
    
    public async init() {
        if (!this.connected) {
            return await this.connect()
        }
        return
    }
    
    private onConnect = () => {
        console.log("Connected")
        this.terminal.reset()
        this.terminal.focus()
        this.attachAddon = new AttachAddon(this.socket!)
        this.terminal.loadAddon(this.attachAddon)
        
        // Handles resize on resize end
        let timeout: number
        timeout = window.setTimeout(this.onResize, 500)
        window.onresize = () => {
            clearTimeout(timeout)
            timeout = window.setTimeout(this.onResize, 500)
        }
    }
    
    private onResize  = () => {
        let cols = this.terminal.cols.toString()
        let rows = this.terminal.rows.toString()
        
        this.socket?.send(`|RESIZE|${cols}:${rows}`)
        try {
            this.fitAddon.fit()
        } catch (e) {
            
        }
    }
    
    private async connect() {
        console.log(`Starting shell session`)
        return new Promise( async (resolve, reject) => {
            try {
                this.socket = new WebSocket(getWsURL("shws/"))
                this.socket.onopen = () => {
                    this.onConnect()
                }
                this.socket.onerror = () => {
                    this.connected = false
                    reject("web socket connection error")
                }
                
                this.pingInterval = window.setInterval(this.ping, 3000)
                this.connected = true
                resolve(true)
                
            } catch {
                clearInterval(this.pingInterval)
                await sleep(1)
                this.connect()
            }
        })
    }
    
    private ping = () => {
        if (this.socket) {
            if ( (this.socket.readyState === 3 )  ) {
                // console.log("Socket not ready")
                clearInterval(this.pingInterval)
                this.connect()
            }
        }
    }
    
    public closeConnection() {
        clearInterval(this.pingInterval)
        this.socket?.close()
        console.log(`Shell session stopped`)
        this.connected = false
    }
}