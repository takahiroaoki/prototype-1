export class ViewModel {
    protected elem: HTMLElement
    protected ref: string
    private emitter = new EventTarget()

    constructor(elem: HTMLElement) {
        this.elem = elem
        this.ref = elem.dataset.ref ?? ''
    }

    protected select(selector: string): HTMLElement | null {
        return this.elem.querySelector(selector)
    }

    protected emit(eventName: string): void {
        this.emitter.dispatchEvent(new CustomEvent(eventName))
    }
}