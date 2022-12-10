package main

import "fmt"

// 命令模式

// 将请求转换为一个包含与请求相关所有信息的独立对象，该转换可以根据不同的请求将方法参数化、延迟执行或撤销

// 方法参数化指的是将每个请求参数传入具体命令工厂创建命令，同时具体命令会默认设置好接受对象，这样可以不管请求参数及类型
// 或者接受对象有几个，都会被封装到具体命令对象的成员字段，并统一通过接口方法调用，屏蔽请求之间的差异

// 例子 电饭煲做饭：根据不同的命令，自动设置不同的参数，用户不用关心，只需要点击开始就ok

// 包含1.具体的命令执行者（电饭煲） 2. 抽象的命令接口 3. 具体的命令 4. 命令触发

type ElectricCooker struct {
	fire     string // 火力
	pressure string // 压力
}

func (e *ElectricCooker) SetFire(fire string) {
	e.fire = fire
}

func (e *ElectricCooker) SetPressure(pressure string) {
	e.pressure = pressure
}

func (e *ElectricCooker) Run(duration string) string {
	return fmt.Sprintf("电饭煲设置火力：%s, 压力：%s, 持续运行：%s\n", e.fire, e.pressure, duration)
}

func (e *ElectricCooker) Shutdown() string {
	return "电饭煲停止运行"
}

// 命令

type CookCommand interface {
	Execute() string // 指令执行方法
}

// steamRiceCommand 蒸饭命令
type steamRiceCommand struct {
	electricCooker *ElectricCooker
}

func NewSteamRiceCommand(electricCooker *ElectricCooker) *steamRiceCommand {
	return &steamRiceCommand{
		electricCooker: electricCooker,
	}
}

func (s *steamRiceCommand) Execute() string {
	s.electricCooker.SetFire("中火")
	s.electricCooker.SetPressure("正常")
	return "蒸饭：" + s.electricCooker.Run("30分钟")
}

// cookCongeeCommand 煮粥命令
type cookCongeeCommand struct {
	electricCooker *ElectricCooker
}

func NewCookCongeeCommand(electricCooker *ElectricCooker) *cookCongeeCommand {
	return &cookCongeeCommand{
		electricCooker: electricCooker,
	}
}

func (c *cookCongeeCommand) Execute() string {
	c.electricCooker.SetFire("大火")
	c.electricCooker.SetPressure("强")
	return "煮粥：" + c.electricCooker.Run("45分钟")
}

type shutdownCommand struct {
	electricCooker *ElectricCooker
}

func NewShutdownCommand(electricCooker *ElectricCooker) *shutdownCommand {
	return &shutdownCommand{
		electricCooker: electricCooker,
	}
}

func (s *shutdownCommand) Execute() string {
	return s.electricCooker.Shutdown()
}

// ElectricCookerInvoker 电饭煲指令触发
type ElectricCookerInvoker struct {
	cookCommand CookCommand
}

func (e *ElectricCookerInvoker) SetCookCommand(cookCommand CookCommand) {
	e.cookCommand = cookCommand
}

func (e *ElectricCookerInvoker) ExecuteCookCommand() string {
	return e.cookCommand.Execute()
}

func main() {
	// 创建命令接受者：电饭煲
	electricCooker := new(ElectricCooker)
	// 创建指令触发器
	electricCookerInvoker := new(ElectricCookerInvoker)

	// 蒸饭
	steamRiceCommand := NewSteamRiceCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(steamRiceCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 煮粥
	cookCongeeCommand := NewCookCongeeCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(cookCongeeCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	//停止
	shutdownCommand := NewShutdownCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(shutdownCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
}
