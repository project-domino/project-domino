/** @module note */

import $ from "jquery";

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
		return $(`<h${this.level}>`).text(this.text);
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
	constructor(text, bold = false, italic = false, underline = false) {
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
	toJSON() {
		const out = {text: this.text};
		if(this.bold) out.bold = true;
		if(this.italic) out.italic = true;
		if(this.underline) out.underline = true;
		return out;
	}
}

class ParagraphNode {
	/**
	 * @param {module:note~TextNode[]} textNodes - The nodes in the paragraph.
	 */
	constructor(textNodes = []) {
		if(!Array.isArray(textNodes))
			throw new TypeError("textNodes must be an Array");
		else if(!textNodes.every(n => n instanceof TextNode))
			throw new TypeError("textNodes must be an Array of TextNodes");

		this.nodes = textNodes;
	}
	static fromElement(element) {
		const nodes = [];
		for(const node of element.childNodes) {
			switch(node.nodeType) {
			case Node.ELEMENT_NODE:
				switch(node.tagName) {
				case "SPAN":
					nodes.push(new TextNode(node.textContent));
					break;
				default:
					throw new TypeError(`unknown tag: ${node.tagName} in ${node}`);
				}
				break;
			case Node.TEXT_NODE:
				nodes.push(new TextNode(node.textContent));
				break;
			default:
				throw new TypeError(`unknown node: ${node}`);
			}
		}
		return new ParagraphNode(nodes);
	}

	render() {
		return $("<p>").append(this.nodes.map(node => node.render()));
	}
}

/**
 * nodesFromElement parses the contents of an element and returns the
 * cooresponding nodes.
 * @param {HTMLElement} e - The element whose contents should be parsed.
 * @return {module:note~Node[]} The parsed nodes.
 */
const nodesFromElement = e => {
	if(e instanceof $)
		e = e[0];
	const out = [];
	for(const node of e.childNodes) {
		switch(node.nodeType) {
		case Node.ELEMENT_NODE:
			switch(node.tagName) {
			case "H1":
				out.push(new HeaderNode(node.textContent, 1));
				break;
			case "H2":
				out.push(new HeaderNode(node.textContent, 2));
				break;
			case "H3":
				out.push(new HeaderNode(node.textContent, 3));
				break;
			case "P":
				out.push(ParagraphNode.fromElement(node));
				break;
			default:
				throw new TypeError(`unknown tag: ${node.tagName} in ${node}`);
			}
			break;
		case Node.TEXT_NODE:
			out.push(new ParagraphNode([
				new TextNode(node.textContent),
			]));
			break;
		default:
			throw new TypeError(`unknown node: ${node}`);
		}
	}
	return out;
};

/**
 * A Note is the thing being edited.
 */
class Note {
	constructor(nodes) {
		if(typeof nodes === "string")
			nodes = JSON.parse(nodes);
		this.nodes = nodes;
	}
	static fromElement(element) {
		return new Note(nodesFromElement(element));
	}

	render() {
		console.debug(this.nodes);
		return this.nodes.map(node => node.render());
	}
	toJSON() {
		return this.nodes;
	}
	update(element) {
		this.nodes = nodesFromElement(element);
	}
}

export {
	Note,
	HeaderNode,
	ImageNode,
	ParagraphNode,
	TextNode,
};
