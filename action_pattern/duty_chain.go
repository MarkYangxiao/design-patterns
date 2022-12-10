package main

import (
	"fmt"
)

// 责任链模式，是一种行为模式，每个处理器都可以对请求进行处理，或者传递给链上下一个处理器

// 模拟登机流程
type BoardingProcessor interface {
	SetNextProcessor(processor BoardingProcessor)
	ProcessFor(passenger *Passenger)
}

// 乘客
type Passenger struct {
	name                  string
	hasBoardingPass       bool // 是否办理登记牌
	hasLuggage            bool //是否有行李需要托运
	isPassIdentityCheck   bool // 是否通过身份校验
	isPassSecurityCheck   bool // 是否通过安检
	isCompleteForBoarding bool // 是否完成登机
}

// baseBoardingProcessor 登机流程处理基类
type baseBoardingProcessor struct {
	// nextProcessor 下一个处理流程
	nextProcessor BoardingProcessor
}

func (b *baseBoardingProcessor) SetNextProcessor(processor BoardingProcessor) {
	b.nextProcessor = processor
}

// ProcessFor 基类中统一实现下一个处理器流转
func (b *baseBoardingProcessor) ProcessFor(passenger *Passenger) {
	if b.nextProcessor != nil {
		b.nextProcessor.ProcessFor(passenger)
	}
}

// boardingPassProcessor 办理登机牌处理器
type boardingPassProcessor struct {
	baseBoardingProcessor // 引用基类
}

func (b *boardingPassProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("为旅客[%s]办理登机牌\n", passenger.name)
		passenger.hasBoardingPass = true
	}

	// 成果办理登机牌进入下一流程
	b.baseBoardingProcessor.ProcessFor(passenger)
}

// luggageCheckInProcessor 办理行李托运处理器
type luggageCheckInProcessor struct {
	baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessFor(passenger *Passenger) {
	// 没办理登机牌不行
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客[%s]未办理登机牌，不能托运行李\n", passenger.name)
		return
	}
	if passenger.hasLuggage {
		fmt.Printf("旅客[%s]办理托运行李\n", passenger.name)
	}
	l.baseBoardingProcessor.ProcessFor(passenger)
}

// identityCheckProcessor 校验身份处理
type identityCheckProcessor struct {
	baseBoardingProcessor
}

func (i *identityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客[%s]未办理登机牌，不能办理身份校验\n", passenger.name)
		return
	}
	if !passenger.isPassIdentityCheck {
		fmt.Printf("旅客[%s]核实身份信息\n", passenger.name)
		passenger.isPassIdentityCheck = true
	}
	i.baseBoardingProcessor.ProcessFor(passenger)
}

// securityCheckProcessor 安检处理器
type securityCheckProcessor struct {
	baseBoardingProcessor
}

func (s *securityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客[%s]未办理登机牌，不能进行安检\n", passenger.name)
		return
	}
	if !passenger.isPassSecurityCheck {
		fmt.Printf("旅客[%s]进行安检\n", passenger.name)
		passenger.isPassSecurityCheck = true
	}
	s.baseBoardingProcessor.ProcessFor(passenger)
}

// completedBoardingProcessor 完成登机处理器
type completedBoardingProcessor struct {
	baseBoardingProcessor
}

func (c *completedBoardingProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass || !passenger.isPassIdentityCheck ||
		!passenger.isPassSecurityCheck {
		fmt.Printf("旅客[%s]登机检查未完成，不能登机\n", passenger.name)
		return
	}
	passenger.isCompleteForBoarding = true
	fmt.Printf("旅客[%s]登机成功\n", passenger.name)
}

func BuildBoardingProcessorChain() BoardingProcessor {
	completedBoardingNode := &completedBoardingProcessor{}

	securityCheckNode := &securityCheckProcessor{}
	securityCheckNode.SetNextProcessor(completedBoardingNode)

	identityCheckNode := &identityCheckProcessor{}
	identityCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckNode := &luggageCheckInProcessor{}
	luggageCheckNode.SetNextProcessor(identityCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckNode)
	return boardingPassNode
}

func main() {
	boardingPassProcessor := BuildBoardingProcessorChain()
	passenger := &Passenger{
		name:                  "zhangsan",
		hasBoardingPass:       false,
		hasLuggage:            true,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}
	boardingPassProcessor.ProcessFor(passenger)
}
