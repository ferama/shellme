
// https://www.pluralsight.com/guides/how-to-communicate-between-independent-components-in-reactjs
export const EventBus = {
    on(event: string, callback: EventListener) {
        document.addEventListener(event, (e: any) => callback(e.detail))
    },
    
    dispatch(event: string, data: EventListener) {
        document.dispatchEvent(new CustomEvent(event, { detail: data }))
    },
    
    remove(event: string, callback: EventListener) {
        document.removeEventListener(event, callback)
    },
}