import test from "node:test"
import assert from "node:assert"
import { DoublyLinkedList, ValueNode } from "./doubly_linked_list.js"

test('just a simple iteration', () => {
	const a = new DoublyLinkedList()
	a.append(new ValueNode('foo'))
	a.append(new ValueNode('bar'))
	a.append(new ValueNode('baz'))

	for (const node of a) {
		console.log('loop', node.value)
	}
})

test('setting the next ValueNode should fail after setting it once', () => {
	const a = new ValueNode('original')
	a.next = new ValueNode('proper')
	a.next = new ValueNode('faulty')
	assert.strictEqual(a.next.value, 'proper')
})

test('setting both next and previous nodes', () => {
	const list = Object.freeze(['foo', 'bar', 'baz', 'tail'])
	const a = new ValueNode('head')
	let temp = a
	for (let i = 0; i < list.length; i++) {
		temp.next = new ValueNode(list[i])
		temp.next.prev = temp
		temp = temp.next
	}

	for (const node of a.next) {
		console.log('loop 2', node.value)
	}

	console.log('head', a.next.next.getHead().value)
})
