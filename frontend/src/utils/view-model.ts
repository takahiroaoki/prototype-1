export class ViewModel {
    private elem: HTMLElement
    private emitter = new EventTarget()

    constructor(elem: HTMLElement) {
        this.elem = elem
    }

    protected select(selector: string): HTMLElement | null {
        return this.elem.querySelector(selector)
    }

    protected emit(eventName: string): void {
        this.emitter.dispatchEvent(new CustomEvent(eventName))
    }
}