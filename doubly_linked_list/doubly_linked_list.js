export class ValueNode {
	/** @type {string} */
	#value

	get value() {
		return this.#value
	}

	/** @type {ValueNode} */
	#next = null

	get next() {
		return this.#next
	}

	set next(node) {
		if (this.#next === null) {
			this.#next = node
		}
	}

	/** @type {ValueNode} */
	#prev = null

	get prev() {
		return this.#prev
	}

	set prev(node) {
		if (this.#prev === null) {
			this.#prev = node
		}
	}

	/**
	 * @param {string} value
	 */
	constructor(value) {
		this.#value = value
	}

	getHead() {
		const recursiveHead = (node) => node.prev !== null ? recursiveHead(node.prev) : node
		return recursiveHead(this)
	}

	*[Symbol.iterator]() {
		let node = this

		while (node !== null) {
			yield node
			node = node.next
		}
	}
}

export class DoublyLinkedList {
	/** @type {ValueNode} */
	#head = null

	/** @type {ValueNode} */
	#tail = null

	/** @param {ValueNode} node */
	constructor(node = null) {
		if (!!node) {
			this.#head = this.#tail = node
		}
	}

	/** @param {ValueNode} node */
	append(node) {
		if (!this.#head) {
			this.#head = node
		} else {
			this.#tail.next = node
		}
		this.#tail = node
	}

	*[Symbol.iterator]() {
		let node = this.#head

		while (node !== null) {
			yield node
			node = node.next
		}
	}
}
