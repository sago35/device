// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from LPC408x_7x_v0.7.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// LPC408x/7x M4
//


.syntax unified

// This is the default handler for interrupts, if triggered but not defined.
.section .text.Default_Handler
.global  Default_Handler
.type    Default_Handler, %function
Default_Handler:
    wfe
    b    Default_Handler
.size Default_Handler, .-Default_Handler

// Avoid the need for repeated .weak and .set instructions.
.macro IRQ handler
    .weak  \handler
    .set   \handler, Default_Handler
.endm

// Must set the "a" flag on the section:
// https://svnweb.freebsd.org/base/stable/11/sys/arm/arm/locore-v4.S?r1=321049&r2=321048&pathrev=321049
// https://sourceware.org/binutils/docs/as/Section.html#ELF-Version
.section .isr_vector, "a", %progbits
.global  __isr_vector
__isr_vector:
    // Interrupt vector as defined by Cortex-M, starting with the stack top.
    // On reset, SP is initialized with *0x0 and PC is loaded with *0x4, loading
    // _stack_top and Reset_Handler.
    .long _stack_top
    .long Reset_Handler
    .long NMI_Handler
    .long HardFault_Handler
    .long MemoryManagement_Handler
    .long BusFault_Handler
    .long UsageFault_Handler
    .long 0
    .long 0
    .long 0
    .long 0
    .long SVC_Handler
    .long DebugMon_Handler
    .long 0
    .long PendSV_Handler
    .long SysTick_Handler

    // Extra interrupts for peripherals defined by the hardware vendor.
    .long WWDT_IRQHandler
    .long TIMER0_IRQHandler
    .long TIMER1_IRQHandler
    .long TIMER2_IRQHandler
    .long TIMER3_IRQHandler
    .long UART0_IRQHandler
    .long UART1_IRQHandler
    .long UART2_IRQHandler
    .long UART3_IRQHandler
    .long PWM1_IRQHandler
    .long I2C0_IRQHandler
    .long I2C1_IRQHandler
    .long I2C2_IRQHandler
    .long 0
    .long SSP0_IRQHandler
    .long SSP1_IRQHandler
    .long 0
    .long RTC_IRQHandler
    .long EINT0_IRQHandler
    .long EINT1_IRQHandler
    .long EINT2_IRQHandler
    .long EINT3_IRQHandler
    .long ADC_IRQHandler
    .long BOD_IRQHandler
    .long USB_IRQHandler
    .long CAN_IRQHandler
    .long GPDMA_IRQHandler
    .long I2S_IRQHandler
    .long ETHERNET_IRQHandler
    .long SDMMC_IRQHandler
    .long MCPWM_IRQHandler
    .long QEI_IRQHandler
    .long 0
    .long USB_NEED_CLK_IRQHandler
    .long 0
    .long UART4_IRQHandler
    .long SSP2_IRQHandler
    .long LCD_IRQHandler
    .long GPIOINT_IRQHandler
    .long PWM0_IRQHandler
    .long EEPROM_IRQHandler
    .long CMP0_IRQHandler
    .long CMP1_IRQHandler

    // Define default implementations for interrupts, redirecting to
    // Default_Handler when not implemented.
    IRQ NMI_Handler
    IRQ HardFault_Handler
    IRQ MemoryManagement_Handler
    IRQ BusFault_Handler
    IRQ UsageFault_Handler
    IRQ SVC_Handler
    IRQ DebugMon_Handler
    IRQ PendSV_Handler
    IRQ SysTick_Handler
    IRQ WWDT_IRQHandler
    IRQ TIMER0_IRQHandler
    IRQ TIMER1_IRQHandler
    IRQ TIMER2_IRQHandler
    IRQ TIMER3_IRQHandler
    IRQ UART0_IRQHandler
    IRQ UART1_IRQHandler
    IRQ UART2_IRQHandler
    IRQ UART3_IRQHandler
    IRQ PWM1_IRQHandler
    IRQ I2C0_IRQHandler
    IRQ I2C1_IRQHandler
    IRQ I2C2_IRQHandler
    IRQ SSP0_IRQHandler
    IRQ SSP1_IRQHandler
    IRQ RTC_IRQHandler
    IRQ EINT0_IRQHandler
    IRQ EINT1_IRQHandler
    IRQ EINT2_IRQHandler
    IRQ EINT3_IRQHandler
    IRQ ADC_IRQHandler
    IRQ BOD_IRQHandler
    IRQ USB_IRQHandler
    IRQ CAN_IRQHandler
    IRQ GPDMA_IRQHandler
    IRQ I2S_IRQHandler
    IRQ ETHERNET_IRQHandler
    IRQ SDMMC_IRQHandler
    IRQ MCPWM_IRQHandler
    IRQ QEI_IRQHandler
    IRQ USB_NEED_CLK_IRQHandler
    IRQ UART4_IRQHandler
    IRQ SSP2_IRQHandler
    IRQ LCD_IRQHandler
    IRQ GPIOINT_IRQHandler
    IRQ PWM0_IRQHandler
    IRQ EEPROM_IRQHandler
    IRQ CMP0_IRQHandler
    IRQ CMP1_IRQHandler
