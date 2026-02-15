/*

One of the services provided by an operating system is memory management. The OS typically provides an API for allocating and releasing memory in a process's address space. A process should only read and write memory at addresses which have been allocated by the operating system. In this kata you will implement a simulation of a simple memory manager.

The language you will be using has no low level memory API, so for our simulation we will simply use an array as the process address space. The memory manager constructor will accept an array (further referred to as memory) where blocks of indices will be allocated later.

Memory Manager Contract
allocate(size)
allocate reserves a sequential block (sub-array) of size received as an argument in memory. It should return the index of the first element in the allocated block, or throw an exception if there is no block big enough to satisfy the requirements.

release(pointer)
release accepts an integer representing the start of the block allocated ealier, and frees that block. If the released block is adjacent to a free block, they should be merged into a larger free block. Releasing an unallocated block should cause an exception.

write(pointer, value)
To support testing this simulation our memory manager needs to support read/write functionality. Only elements within allocated blocks may be written to. The write method accepts an index in memory and a value. The value should be stored in memory at that index if it lies within an allocated block, or throw an exception otherwise.

read(pointer)
This method is the counterpart to write. Only indices within allocated blocks may be read. The read method accepts an index. If the index is within an allocated block, the value stored in memory at that index should be returned, or an exception should be thrown otherwise.

*/

class MemoryManager {
    /**
     * @constructor Creates a new memory manager for the provided array.
     * @param {array} memory An array to use as the backing memory.
     */
    constructor(memory) {
        this.memory = memory;
        this.freeBlocks = [{start: 0, size: memory.length}];
        this.allocatedBlocks = new Map();
    }

    /**
     * Allocates a block of memory of requested size.
     * @param {number} size - The size of the block to allocate.
     * @returns {number} A pointer which is the index of the first location in the allocated block.
     * @throws If it is not possible to allocate a block of the requested size.
     */
    allocate(size) {
        if (size <= 0) {
            throw new Error('Size must be positive');
        }

        for (let i = 0; i < this.freeBlocks.length; i++) {
            const freeBlock = this.freeBlocks[i];

            if (freeBlock.size >= size) {
                const pointer = freeBlock.start;
                this.allocatedBlocks.set(pointer, {start: pointer, size: size});

                if (freeBlock.size === size) {
                    this.freeBlocks.splice(i, 1);
                } else {
                    freeBlock.start += size;
                    freeBlock.size -= size;
                }

                return pointer;
            }
        }

        throw new Error('Not enough memory to allocate block of size ' + size);
    }

    /**
     * Releases a previously allocated block of memory.
     * @param {number} pointer - The pointer to the block to release.
     * @throws If the pointer does not point to an allocated block.
     */
    release(pointer) {
        const allocatedBlock = this.allocatedBlocks.get(pointer);

        if (!allocatedBlock) {
            throw new Error('Pointer does not point to an allocated block');
        }

        this.allocatedBlocks.delete(pointer);
        const newFreeBlock = {start: pointer, size: allocatedBlock.size};

        let inserted = false;
        for (let i = 0; i < this.freeBlocks.length; i++) {
            if (newFreeBlock.start < this.freeBlocks[i].start) {
                this.freeBlocks.splice(i, 0, newFreeBlock);
                inserted = true;
                break;
            }
        }

        if (!inserted) {
            this.freeBlocks.push(newFreeBlock);
        }

        // Merge free blocks
        for (let i = 0; i < this.freeBlocks.length - 1; i++) {
            const current = this.freeBlocks[i];
            const next = this.freeBlocks[i + 1];

            if (current.start + current.size === next.start) {
                current.size += next.size;
                this.freeBlocks.splice(i + 1, 1);
                i--;
            }
        }
    }

    /**
     * Reads the value at the location identified by pointer
     * @param {number} pointer - The location to read.
     * @returns {number} The value at that location.
     * @throws If pointer is in unallocated memory.
     */
    read(pointer) {
        if (!this.isAllocated(pointer)) {
            throw new Error('Pointer is in unallocated memory');
        }

        return this.memory[pointer];
    }

    /**
     * Writes a value to the location identified by pointer
     * @param {number} pointer - The location to write to.
     * @param {number} value - The value to write.
     * @throws If pointer is in unallocated memory.
     */
    write(pointer, value) {
        if (!this.isAllocated(pointer)) {
            throw new Error('Pointer is in unallocated memory');
        }

        this.memory[pointer] = value;
    }

    /**
     * Checks if a pointer is within an allocated block
     * @param {number} pointer - The pointer to check
     * @returns {boolean} True if the pointer is in allocated memory
     */
    isAllocated(pointer) {
        for (const [_, block] of this.allocatedBlocks) {
            if (pointer >= block.start && pointer < block.start + block.size) {
                return true;
            }
        }
        return false;
    }
}
