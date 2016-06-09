export default class EventEmitter {
	constructor() {
		this.listeners = new Map();
	}
	emit(event, ...args) {
		const callbacks = this.listeners.get(event);
		if(callbacks)
			callbacks.forEach(cb => cb(...args));
		else
			console.warn(`no listeners for ${event}...`);
	}
	on(event, cb) {
		if(!this.listeners.has(event))
			this.listeners.set(event, []);
		this.listeners.get(event).push(cb);
	}
}
