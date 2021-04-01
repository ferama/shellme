export interface Cell {
    id: number,
    code: string,
    status: string,
    runID: number
}

export interface LogItem {
    runID: number,
    cellID: number,
    data: string,
    channel: string,
    contentType: string
}

export interface Document {
    cells: Cell[]
    filename: string
    logs: { [cellID: number]: LogItem[] }
}