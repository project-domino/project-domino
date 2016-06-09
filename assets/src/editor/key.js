export default class Key {
	constructor(spec) {
		if(typeof spec === "string") {
			const parts = spec.split("-");
			this.modifiers = parts.slice(0, parts.length - 1)
				.map(ch => {
					switch(ch.toUpperCase()) {
					case "C":
						return "ctrl";
					case "A":
						return "alt";
					case "S":
						return "shift";
					default:
						throw new Error(`Unknown Key modifier: ${ch}`);
					}
				});
			this.rawKey = parts[parts.length - 1].toLowerCase();
		} else if(spec instanceof Event) {
			if(spec.code.match(/Digit[0-9]/))
				this.rawKey = spec.code[5];
			else
				this.rawKey = spec.key.toLowerCase();
			this.key = spec.key;

			const modifiers = [];
			if(spec.ctrlKey) modifiers.push("ctrl");
			if(spec.altKey) modifiers.push("alt");
			if(spec.shiftKey) modifiers.push("shift");
			this.modifiers = modifiers;
		} else {
			throw new Error(`Invalid argument to Key: ${spec}`);
		}
	}

	get ctrl() { return this.hasModifier("ctrl"); }
	get alt() { return this.hasModifier("alt"); }
	get shift() { return this.hasModifier("shift"); }

	equals(key) {
		if(this.rawKey !== key.rawKey)
			return false;
		return ["ctrl", "alt", "shift"].every(mod => {
			return this.hasModifier(mod) === key.hasModifier(mod);
		});
	}
	hasModifier(modifier) {
		modifier = modifier.toLowerCase();

		return this.modifiers
			.map(mod => mod.slice(0, modifier.length))
			.map(mod => mod.toLowerCase())
			.filter(mod => mod === modifier)
			.length > 0;
	}
	toString() {
		const modifiers = this.modifiers
			.map(modifier => modifier.charAt(0))
			.map(ch => ch.toUpperCase())
			.map(ch => `${ch}-`)
			.reduce((a, b) => a + b, "");
		return modifiers + this.rawKey;
	}
}
