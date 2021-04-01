export interface WSMessage {
    cellID: number,
    runID: number
    data: string,
    channel: string,
    contentType: string
}