import { Views } from '../enums/views';
import { writable } from 'svelte/store';

export class ViewManager {
    static _stack: Views[] = []
    static view = writable(Views.UPLOAD_FILE)

    static goTo(view: Views) {
        this._stack = [...this._stack, view]
        this.view.set(view)
    }

    static goBack() {
        const view = this._stack.pop()
        this.view.set(view)
    }
}