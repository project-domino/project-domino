export default class AutosaveManager {
	constructor(parent, timeMax = 5000, charMax = 50) {
		this.parent = parent;
		this.charMax = charMax;
		this.timeMax = timeMax;

		this.timer = 0;
		this.chars = 0;

		this.parent.on("keypress", () => this.keypressHandler());
	}
	keypressHandler() {
		clearTimeout(this.timer);
		if(this.chars++ > this.charMax)
			this.save();
		else
			this.timer = setTimeout(() => this.save(), this.timeMax);
	}
	save() {
		this.chars = 0;
		this.timer = 0;

		this.parent.emit("save-triggered");
	}
}
