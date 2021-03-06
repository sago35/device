// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from LPC5410x_v0.4.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// LPC5410x Cortex-M4 MCU; Cortex-M0+ coprocessor
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
    .long WDT_IRQHandler
    .long BOD_IRQHandler
    .long 0
    .long DMA_IRQHandler
    .long GINT0_IRQHandler
    .long PIN_INT0_IRQHandler
    .long PIN_INT1_IRQHandler
    .long PIN_INT2_IRQHandler
    .long PIN_INT3_IRQHandler
    .long UTICK_IRQHandler
    .long MRT_IRQHandler
    .long CT32B0_IRQHandler
    .long CT32B1_IRQHandler
    .long CT32B2_IRQHandler
    .long CT32B3_IRQHandler
    .long CT32B4_IRQHandler
    .long SCT0_IRQHandler
    .long UART0_IRQHandler
    .long UART1_IRQHandler
    .long UART2_IRQHandler
    .long UART3_IRQHandler
    .long I2C0_IRQHandler
    .long I2C1_IRQHandler
    .long I2C2_IRQHandler
    .long SPI0_IRQHandler
    .long SPI1_IRQHandler
    .long ADC_SEQA_IRQHandler
    .long ADC_SEQB_IRQHandler
    .long ADC_THCMP_IRQHandler
    .long RTC_IRQHandler
    .long 0
    .long MAILBOX_IRQHandler
    .long GINT1_IRQHandler
    .long PIN_INT4_IRQHandler
    .long PIN_INT5_IRQHandler
    .long PIN_INT6_IRQHandler
    .long PIN_INT7_IRQHandler
    .long 0
    .long 0
    .long 0
    .long RIT_IRQHandler

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
    IRQ WDT_IRQHandler
    IRQ BOD_IRQHandler
    IRQ DMA_IRQHandler
    IRQ GINT0_IRQHandler
    IRQ PIN_INT0_IRQHandler
    IRQ PIN_INT1_IRQHandler
    IRQ PIN_INT2_IRQHandler
    IRQ PIN_INT3_IRQHandler
    IRQ UTICK_IRQHandler
    IRQ MRT_IRQHandler
    IRQ CT32B0_IRQHandler
    IRQ CT32B1_IRQHandler
    IRQ CT32B2_IRQHandler
    IRQ CT32B3_IRQHandler
    IRQ CT32B4_IRQHandler
    IRQ SCT0_IRQHandler
    IRQ UART0_IRQHandler
    IRQ UART1_IRQHandler
    IRQ UART2_IRQHandler
    IRQ UART3_IRQHandler
    IRQ I2C0_IRQHandler
    IRQ I2C1_IRQHandler
    IRQ I2C2_IRQHandler
    IRQ SPI0_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ ADC_SEQA_IRQHandler
    IRQ ADC_SEQB_IRQHandler
    IRQ ADC_THCMP_IRQHandler
    IRQ RTC_IRQHandler
    IRQ MAILBOX_IRQHandler
    IRQ GINT1_IRQHandler
    IRQ PIN_INT4_IRQHandler
    IRQ PIN_INT5_IRQHandler
    IRQ PIN_INT6_IRQHandler
    IRQ PIN_INT7_IRQHandler
    IRQ RIT_IRQHandler
