function compressNodes(nodes) {
	if(nodes.length === 0)
		return [];
	return nodes.slice(1).reduce((array, node) => {
		const last = array[array.length - 1];
		if(last.mergable(node))
			array[array.length - 1] = last.merge(node);
		else
			array.push(node);
		return array;
	}, [nodes[0]]);
}

class Node {
	constructor(text, formatting) {
		this.text = text;
		this.formatting = new Set(formatting);
	}
	mergable(other) {
		if(this.formatting.size !== other.formatting.size)
			return false;
		let equals = true;
		this.formatting.forEach(f => {
			if(!other.formatting.has(f))
				equals = false;
		});
		return equals;
	}
	merge(other) {
		if(!this.mergable(other))
			return null;
		return new Node(this.text + other.text, this.formatting);
	}
	render() {
		const e = document.createElement("span");
		e.textContent = this.text;
		this.formatting.forEach(cls => e.classList.add(cls));
		return e;
	}
}

export class Document {
	constructor(parent, nodes = []) {
		this.currentFormatting = [];
		this.parent = parent;
		this.nodes = nodes;

		this.parent.on("formatting", type => this.formattingHandler(type));
		this.parent.on("keypress", event => this.keypressHandler(event));
	}
	formattingHandler(type) {
		if(type === "clear")
			this.currentFormatting = [];
		else if(this.currentFormatting.indexOf(type) === -1)
			this.currentFormatting.push(type);
		else
			this.currentFormatting = this.currentFormatting.filter(t => t !== type);
	}
	keypressHandler(event) {
		this.nodes.push(new Node(event.key, this.currentFormatting));
		this.parent.emit("render", this.render());
	}
	render() {
		return compressNodes(this.nodes)
			.map(node => node.render())
			.reduce((a, e) => {
				a.push(e);
				return a;
			}, []);
	}
	toJSON() {
		return compressNodes(this.nodes);
	}
}
