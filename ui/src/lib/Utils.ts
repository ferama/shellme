
export const isDev = ():boolean => {
    const development: boolean = !process.env.NODE_ENV || process.env.NODE_ENV === 'development'
    return development
}

export const sleep = async (secs: number): Promise<boolean> => {
    return new Promise( (resolve, _) => {
        setTimeout(() => {
            resolve(true)
        }, secs * 1000)
    })
}

export const getWsURL = (path: string): string => {

    const proto = ((window.location.protocol === "https:") ? "wss://" : "ws://") 
    const host = window.location.hostname
    let port = window.location.port
    
    if (isDev()) {
        port = "8000"
    }

    return `${proto}${host}:${port}/${path}`
}