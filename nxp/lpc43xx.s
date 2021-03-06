// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from LPC43xx_43Sxx.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// Register cmsis file for LPC43xx parts
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
    .long DAC_IRQHandler
    .long 0
    .long DMA_IRQHandler
    .long 0
    .long FLASH_IRQHandler
    .long ETHERNET_IRQHandler
    .long SDIO_IRQHandler
    .long LCD_IRQHandler
    .long USB0_IRQHandler
    .long USB1_IRQHandler
    .long SCT_IRQHandler
    .long RITIMER_IRQHandler
    .long TIMER0_IRQHandler
    .long TIMER1_IRQHandler
    .long TIMER2_IRQHandler
    .long TIMER3_IRQHandler
    .long MCPWM_IRQHandler
    .long ADC0_IRQHandler
    .long I2C0_IRQHandler
    .long I2C1_IRQHandler
    .long SPI_INT_IRQHandler
    .long ADC1_IRQHandler
    .long SSP0_IRQHandler
    .long SSP1_IRQHandler
    .long USART0_IRQHandler
    .long UART1_IRQHandler
    .long USART2_IRQHandler
    .long USART3_IRQHandler
    .long I2S0_IRQHandler
    .long I2S1_IRQHandler
    .long SPIFI_IRQHandler
    .long SGPIO_IINT_IRQHandler
    .long PIN_INT0_IRQHandler
    .long PIN_INT1_IRQHandler
    .long PIN_INT2_IRQHandler
    .long PIN_INT3_IRQHandler
    .long PIN_INT4_IRQHandler
    .long PIN_INT5_IRQHandler
    .long PIN_INT6_IRQHandler
    .long PIN_INT7_IRQHandler
    .long GINT0_IRQHandler
    .long GINT1_IRQHandler
    .long EVENTROUTER_IRQHandler
    .long C_CAN1_IRQHandler
    .long 0
    .long ADCHS_IRQHandler
    .long ATIMER_IRQHandler
    .long RTC_IRQHandler
    .long 0
    .long WWDT_IRQHandler
    .long 0
    .long C_CAN0_IRQHandler
    .long QEI_IRQHandler

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
    IRQ DAC_IRQHandler
    IRQ DMA_IRQHandler
    IRQ FLASH_IRQHandler
    IRQ ETHERNET_IRQHandler
    IRQ SDIO_IRQHandler
    IRQ LCD_IRQHandler
    IRQ USB0_IRQHandler
    IRQ USB1_IRQHandler
    IRQ SCT_IRQHandler
    IRQ RITIMER_IRQHandler
    IRQ TIMER0_IRQHandler
    IRQ TIMER1_IRQHandler
    IRQ TIMER2_IRQHandler
    IRQ TIMER3_IRQHandler
    IRQ MCPWM_IRQHandler
    IRQ ADC0_IRQHandler
    IRQ I2C0_IRQHandler
    IRQ I2C1_IRQHandler
    IRQ SPI_INT_IRQHandler
    IRQ ADC1_IRQHandler
    IRQ SSP0_IRQHandler
    IRQ SSP1_IRQHandler
    IRQ USART0_IRQHandler
    IRQ UART1_IRQHandler
    IRQ USART2_IRQHandler
    IRQ USART3_IRQHandler
    IRQ I2S0_IRQHandler
    IRQ I2S1_IRQHandler
    IRQ SPIFI_IRQHandler
    IRQ SGPIO_IINT_IRQHandler
    IRQ PIN_INT0_IRQHandler
    IRQ PIN_INT1_IRQHandler
    IRQ PIN_INT2_IRQHandler
    IRQ PIN_INT3_IRQHandler
    IRQ PIN_INT4_IRQHandler
    IRQ PIN_INT5_IRQHandler
    IRQ PIN_INT6_IRQHandler
    IRQ PIN_INT7_IRQHandler
    IRQ GINT0_IRQHandler
    IRQ GINT1_IRQHandler
    IRQ EVENTROUTER_IRQHandler
    IRQ C_CAN1_IRQHandler
    IRQ ADCHS_IRQHandler
    IRQ ATIMER_IRQHandler
    IRQ RTC_IRQHandler
    IRQ WWDT_IRQHandler
    IRQ C_CAN0_IRQHandler
    IRQ QEI_IRQHandler
