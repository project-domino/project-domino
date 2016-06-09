/** @module note */

/**
 * A Note is the thing being edited.
 */
class Note {
	constructor(nodes) {
		this.nodes = nodes;
	}
	// TODO
	// static fromJSON(json) {
	// 	throw new Error("NYI");
	// }

	render() {
		return this.nodes.map(node => node.render());
	}
	toJSON() {
		return JSON.stringify(this.nodes);
	}
}

class HeaderNode {
	constructor(text, level = 1) {
		if(typeof text !== "string")
			throw new TypeError("text must be a string");
		else if(typeof level !== "number")
			throw new TypeError("level must be a number");
		else if(level !== 1 && level !== 2 && level !== 3)
			throw new TypeError("level must be between an integer 1 and 3");

		this.text = text;
		this.level = level;
	}
	render() {
		return $(`<h${this.level + 1}>`).text(this.text);
	}
}

class ImageNode {
	constructor(url) {
		this.url = url;
	}
	render() {
		return $("img").attr("src", this.url);
	}
}

class TextNode {
	/**
	 * @param {string} text - The text in the node.
	 * @param {boolean} bold - Whether the text should be bold.
	 * @param {boolean} italic - Whether the text should be italic.
	 * @param {boolean} underline - Whether the text should be underlined.
	 */
	constructor(text, bold, italic, underline) {
		this.text = text;
		this.bold = bold;
		this.italic = italic;
		this.underline = underline;
	}
	render() {
		let e = $("<span>").text(this.text);
		if(this.bold) e.css("font-weight", "bold");
		if(this.italic) e.css("font-style", "italic");
		if(this.underline) e.css("text-decoration", "underline");
		return e;
	}
}

class ParagraphNode {
	/**
	 * @param {module:note~TextNode[]} textNodes - The nodes in the paragraph.
	 */
	constructor(textNodes) {
		if(!Array.isArray(textNodes))
			throw new TypeError("textNodes must be an Array");
		else if(!textNodes.every(n => n instanceof TextNode))
			throw new TypeError("textNodes must be an Array of TextNodes");

		this.nodes = textNodes;
	}
	render() {
		return $("<p>").append(this.nodes.map(node => node.render()));
	}
}

export {
	Note,
	HeaderNode,
	ImageNode,
	ParagraphNode,
	TextNode,
};
