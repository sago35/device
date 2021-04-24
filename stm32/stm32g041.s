// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from stm32g041.svd, see https://github.com/tinygo-org/stm32-svd

// STM32G041
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
    .long WWDG_IRQHandler
    .long PVD_IRQHandler
    .long RTC_TAMP_IRQHandler
    .long FLASH_IRQHandler
    .long RCC_IRQHandler
    .long EXTI0_1_IRQHandler
    .long EXTI2_3_IRQHandler
    .long EXTI4_15_IRQHandler
    .long 0
    .long DMA_Channel1_IRQHandler
    .long DMA_Channel2_3_IRQHandler
    .long DMA_Channel4_5_6_7_IRQHandler
    .long ADC_COMP_IRQHandler
    .long TIM1_BRK_UP_TRG_COM_IRQHandler
    .long TIM1_CC_IRQHandler
    .long TIM2_IRQHandler
    .long TIM3_IRQHandler
    .long 0
    .long 0
    .long TIM14_IRQHandler
    .long 0
    .long TIM16_IRQHandler
    .long TIM17_IRQHandler
    .long I2C1_IRQHandler
    .long I2C2_IRQHandler
    .long SPI1_IRQHandler
    .long SPI2_IRQHandler
    .long USART1_IRQHandler
    .long USART2_IRQHandler
    .long USART3_USART4_LPUART1_IRQHandler
    .long CEC_IRQHandler
    .long AES_RNG_IRQHandler

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
    IRQ WWDG_IRQHandler
    IRQ PVD_IRQHandler
    IRQ RTC_TAMP_IRQHandler
    IRQ FLASH_IRQHandler
    IRQ RCC_IRQHandler
    IRQ EXTI0_1_IRQHandler
    IRQ EXTI2_3_IRQHandler
    IRQ EXTI4_15_IRQHandler
    IRQ DMA_Channel1_IRQHandler
    IRQ DMA_Channel2_3_IRQHandler
    IRQ DMA_Channel4_5_6_7_IRQHandler
    IRQ ADC_COMP_IRQHandler
    IRQ TIM1_BRK_UP_TRG_COM_IRQHandler
    IRQ TIM1_CC_IRQHandler
    IRQ TIM2_IRQHandler
    IRQ TIM3_IRQHandler
    IRQ TIM14_IRQHandler
    IRQ TIM16_IRQHandler
    IRQ TIM17_IRQHandler
    IRQ I2C1_IRQHandler
    IRQ I2C2_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ SPI2_IRQHandler
    IRQ USART1_IRQHandler
    IRQ USART2_IRQHandler
    IRQ USART3_USART4_LPUART1_IRQHandler
    IRQ CEC_IRQHandler
    IRQ AES_RNG_IRQHandler
