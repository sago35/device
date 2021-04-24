// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from ATSAME70J21.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/Atmel

// Microchip ATSAME70J21 Microcontroller
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
    .long SUPC_IRQHandler
    .long RSTC_IRQHandler
    .long RTC_IRQHandler
    .long RTT_IRQHandler
    .long WDT_IRQHandler
    .long PMC_IRQHandler
    .long EFC_IRQHandler
    .long UART0_IRQHandler
    .long UART1_IRQHandler
    .long 0
    .long PIOA_IRQHandler
    .long PIOB_IRQHandler
    .long 0
    .long USART0_IRQHandler
    .long USART1_IRQHandler
    .long 0
    .long PIOD_IRQHandler
    .long 0
    .long 0
    .long TWIHS0_IRQHandler
    .long TWIHS1_IRQHandler
    .long 0
    .long SSC_IRQHandler
    .long TC0_IRQHandler
    .long TC1_IRQHandler
    .long TC2_IRQHandler
    .long TC3_IRQHandler
    .long TC4_IRQHandler
    .long TC5_IRQHandler
    .long AFEC0_IRQHandler
    .long DACC_IRQHandler
    .long PWM0_IRQHandler
    .long ICM_IRQHandler
    .long ACC_IRQHandler
    .long USBHS_IRQHandler
    .long MCAN0_INT0_IRQHandler
    .long MCAN0_INT1_IRQHandler
    .long 0
    .long 0
    .long GMAC_IRQHandler
    .long AFEC1_IRQHandler
    .long 0
    .long 0
    .long QSPI_IRQHandler
    .long UART2_IRQHandler
    .long 0
    .long 0
    .long TC6_IRQHandler
    .long TC7_IRQHandler
    .long TC8_IRQHandler
    .long TC9_IRQHandler
    .long TC10_IRQHandler
    .long TC11_IRQHandler
    .long 0
    .long 0
    .long 0
    .long AES_IRQHandler
    .long TRNG_IRQHandler
    .long XDMAC_IRQHandler
    .long ISI_IRQHandler
    .long PWM1_IRQHandler
    .long FPU_IRQHandler
    .long 0
    .long RSWDT_IRQHandler
    .long 0
    .long 0
    .long GMAC_Q1_IRQHandler
    .long GMAC_Q2_IRQHandler
    .long IXC_IRQHandler

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
    IRQ SUPC_IRQHandler
    IRQ RSTC_IRQHandler
    IRQ RTC_IRQHandler
    IRQ RTT_IRQHandler
    IRQ WDT_IRQHandler
    IRQ PMC_IRQHandler
    IRQ EFC_IRQHandler
    IRQ UART0_IRQHandler
    IRQ UART1_IRQHandler
    IRQ PIOA_IRQHandler
    IRQ PIOB_IRQHandler
    IRQ USART0_IRQHandler
    IRQ USART1_IRQHandler
    IRQ PIOD_IRQHandler
    IRQ TWIHS0_IRQHandler
    IRQ TWIHS1_IRQHandler
    IRQ SSC_IRQHandler
    IRQ TC0_IRQHandler
    IRQ TC1_IRQHandler
    IRQ TC2_IRQHandler
    IRQ TC3_IRQHandler
    IRQ TC4_IRQHandler
    IRQ TC5_IRQHandler
    IRQ AFEC0_IRQHandler
    IRQ DACC_IRQHandler
    IRQ PWM0_IRQHandler
    IRQ ICM_IRQHandler
    IRQ ACC_IRQHandler
    IRQ USBHS_IRQHandler
    IRQ MCAN0_INT0_IRQHandler
    IRQ MCAN0_INT1_IRQHandler
    IRQ GMAC_IRQHandler
    IRQ AFEC1_IRQHandler
    IRQ QSPI_IRQHandler
    IRQ UART2_IRQHandler
    IRQ TC6_IRQHandler
    IRQ TC7_IRQHandler
    IRQ TC8_IRQHandler
    IRQ TC9_IRQHandler
    IRQ TC10_IRQHandler
    IRQ TC11_IRQHandler
    IRQ AES_IRQHandler
    IRQ TRNG_IRQHandler
    IRQ XDMAC_IRQHandler
    IRQ ISI_IRQHandler
    IRQ PWM1_IRQHandler
    IRQ FPU_IRQHandler
    IRQ RSWDT_IRQHandler
    IRQ GMAC_Q1_IRQHandler
    IRQ GMAC_Q2_IRQHandler
    IRQ IXC_IRQHandler
