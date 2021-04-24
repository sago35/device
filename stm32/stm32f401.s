// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from stm32f401.svd, see https://github.com/tinygo-org/stm32-svd

// STM32F401
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
    .long 0
    .long PVD_IRQHandler
    .long TAMP_STAMP_IRQHandler
    .long RTC_WKUP_IRQHandler
    .long FLASH_IRQHandler
    .long RCC_IRQHandler
    .long EXTI0_IRQHandler
    .long EXTI1_IRQHandler
    .long EXTI2_IRQHandler
    .long EXTI3_IRQHandler
    .long EXTI4_IRQHandler
    .long DMA1_STREAM0_IRQHandler
    .long DMA1_STREAM1_IRQHandler
    .long DMA1_STREAM2_IRQHandler
    .long DMA1_STREAM3_IRQHandler
    .long DMA1_STREAM4_IRQHandler
    .long DMA1_STREAM5_IRQHandler
    .long DMA1_STREAM6_IRQHandler
    .long ADC_IRQHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long EXTI9_5_IRQHandler
    .long TIM1_BRK_TIM9_IRQHandler
    .long TIM1_UP_TIM10_IRQHandler
    .long TIM1_TRG_COM_TIM11_IRQHandler
    .long TIM1_CC_IRQHandler
    .long TIM2_IRQHandler
    .long TIM3_IRQHandler
    .long TIM4_IRQHandler
    .long I2C1_EV_IRQHandler
    .long I2C1_ER_IRQHandler
    .long I2C2_EV_IRQHandler
    .long I2C2_ER_IRQHandler
    .long SPI1_IRQHandler
    .long SPI2_IRQHandler
    .long USART1_IRQHandler
    .long USART2_IRQHandler
    .long 0
    .long EXTI15_10_IRQHandler
    .long RTC_Alarm_IRQHandler
    .long OTG_FS_WKUP_IRQHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long DMA1_STREAM7_IRQHandler
    .long 0
    .long SDIO_IRQHandler
    .long TIM5_IRQHandler
    .long SPI3_IRQHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long DMA2_STREAM0_IRQHandler
    .long DMA2_STREAM1_IRQHandler
    .long DMA2_STREAM2_IRQHandler
    .long DMA2_STREAM3_IRQHandler
    .long DMA2_STREAM4_IRQHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long OTG_FS_IRQHandler
    .long DMA2_STREAM5_IRQHandler
    .long DMA2_STREAM6_IRQHandler
    .long DMA2_STREAM7_IRQHandler
    .long USART6_IRQHandler
    .long I2C3_EV_IRQHandler
    .long I2C3_ER_IRQHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long FPU_IRQHandler
    .long 0
    .long 0
    .long SPI4_IRQHandler

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
    IRQ PVD_IRQHandler
    IRQ TAMP_STAMP_IRQHandler
    IRQ RTC_WKUP_IRQHandler
    IRQ FLASH_IRQHandler
    IRQ RCC_IRQHandler
    IRQ EXTI0_IRQHandler
    IRQ EXTI1_IRQHandler
    IRQ EXTI2_IRQHandler
    IRQ EXTI3_IRQHandler
    IRQ EXTI4_IRQHandler
    IRQ DMA1_STREAM0_IRQHandler
    IRQ DMA1_STREAM1_IRQHandler
    IRQ DMA1_STREAM2_IRQHandler
    IRQ DMA1_STREAM3_IRQHandler
    IRQ DMA1_STREAM4_IRQHandler
    IRQ DMA1_STREAM5_IRQHandler
    IRQ DMA1_STREAM6_IRQHandler
    IRQ ADC_IRQHandler
    IRQ EXTI9_5_IRQHandler
    IRQ TIM1_BRK_TIM9_IRQHandler
    IRQ TIM1_UP_TIM10_IRQHandler
    IRQ TIM1_TRG_COM_TIM11_IRQHandler
    IRQ TIM1_CC_IRQHandler
    IRQ TIM2_IRQHandler
    IRQ TIM3_IRQHandler
    IRQ TIM4_IRQHandler
    IRQ I2C1_EV_IRQHandler
    IRQ I2C1_ER_IRQHandler
    IRQ I2C2_EV_IRQHandler
    IRQ I2C2_ER_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ SPI2_IRQHandler
    IRQ USART1_IRQHandler
    IRQ USART2_IRQHandler
    IRQ EXTI15_10_IRQHandler
    IRQ RTC_Alarm_IRQHandler
    IRQ OTG_FS_WKUP_IRQHandler
    IRQ DMA1_STREAM7_IRQHandler
    IRQ SDIO_IRQHandler
    IRQ TIM5_IRQHandler
    IRQ SPI3_IRQHandler
    IRQ DMA2_STREAM0_IRQHandler
    IRQ DMA2_STREAM1_IRQHandler
    IRQ DMA2_STREAM2_IRQHandler
    IRQ DMA2_STREAM3_IRQHandler
    IRQ DMA2_STREAM4_IRQHandler
    IRQ OTG_FS_IRQHandler
    IRQ DMA2_STREAM5_IRQHandler
    IRQ DMA2_STREAM6_IRQHandler
    IRQ DMA2_STREAM7_IRQHandler
    IRQ USART6_IRQHandler
    IRQ I2C3_EV_IRQHandler
    IRQ I2C3_ER_IRQHandler
    IRQ FPU_IRQHandler
    IRQ SPI4_IRQHandler