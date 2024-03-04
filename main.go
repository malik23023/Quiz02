package main

import (
    "fmt"
)

type MyBlock struct {
    Data string
}

type MyBlockchain struct {
    Blocks []*MyBlock
}

func (bc *MyBlockchain) PrintBlocks() {
    fmt.Println("Blocks in the blockchain:")
    for _, block := range bc.Blocks {
        fmt.Println(block.Data)
    }
}

func (bc *MyBlockchain) AddNewBlock(data string) {
    newBlock := &MyBlock{Data: data}
    bc.Blocks = append(bc.Blocks, newBlock)
    fmt.Println("New block added to the blockchain with data:", data)
}

func (bc *MyBlockchain) ModifyBlockData(index int, newData string) {
    if index >= 0 && index < len(bc.Blocks) {
        bc.Blocks[index].Data = newData
        fmt.Println("Block at index", index, "modified with new data:", newData)
    } else {
        fmt.Println("Invalid index")
    }
}

func main() {
    blockchain := &MyBlockchain{}

    blockchain.AddNewBlock("Block 1")
    blockchain.AddNewBlock("Block 2")
    blockchain.AddNewBlock("Block 3")
    blockchain.PrintBlocks()
    blockchain.ModifyBlockData(1, "Modified Block 2")
    blockchain.PrintBlocks()
}

