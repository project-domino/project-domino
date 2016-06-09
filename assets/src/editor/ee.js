/** @module ee */

/**
 * An EventEmitter emits and handles events.
 */
class EventEmitter {
	/**
	 * Constructs a new EventEmitter.
	 */
	constructor() {
		this.listeners = new Map();
	}

	/**
	 * Sends an event.
	 * @param {string} event - The event to send.
	 * @param ...args - Any other data to be sent.
	 * @return {module:ee~EventEmitter} Itself, to allow for chaining.
	 */
	emit(event, ...args) {
		const callbacks = this.listeners.get(event);
		if(callbacks)
			callbacks.forEach(cb => cb(...args));
		else
			console.warn(`no listeners for ${event}...`);
		return this;
	}

	/**
	 * Registers an event handler.
	 * @param {string} event - The event to handle.
	 * @param {module:ee~EventEmitter~handler} cb - The handler.
	 * @return {module:ee~EventEmitter} Itself, to allow for chaining.
	 */
	on(event, cb) {
		if(!this.listeners.has(event))
			this.listeners.set(event, []);
		this.listeners.get(event).push(cb);
		return this;
	}
}

/**
 * A callback to handle emitted events.
 * @callback module:ee~EventEmitter~handler
 * @param {string} event - The event that was emitted.
 * @param ...args - Any other data that was sent.
 */

export default EventEmitter;
