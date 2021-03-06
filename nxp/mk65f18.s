// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from MK65F18.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// MK65F18 NXP Microcontroller
//
//     Copyright 2016-2018 NXP SPDX-License-Identifier: BSD-3-Clause

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
    .long DMA0_DMA16_IRQHandler
    .long DMA1_DMA17_IRQHandler
    .long DMA2_DMA18_IRQHandler
    .long DMA3_DMA19_IRQHandler
    .long DMA4_DMA20_IRQHandler
    .long DMA5_DMA21_IRQHandler
    .long DMA6_DMA22_IRQHandler
    .long DMA7_DMA23_IRQHandler
    .long DMA8_DMA24_IRQHandler
    .long DMA9_DMA25_IRQHandler
    .long DMA10_DMA26_IRQHandler
    .long DMA11_DMA27_IRQHandler
    .long DMA12_DMA28_IRQHandler
    .long DMA13_DMA29_IRQHandler
    .long DMA14_DMA30_IRQHandler
    .long DMA15_DMA31_IRQHandler
    .long DMA_Error_IRQHandler
    .long MCM_IRQHandler
    .long FTFE_IRQHandler
    .long Read_Collision_IRQHandler
    .long LVD_LVW_IRQHandler
    .long LLWU_IRQHandler
    .long WDOG_EWM_IRQHandler
    .long RNG_IRQHandler
    .long I2C0_IRQHandler
    .long I2C1_IRQHandler
    .long SPI0_IRQHandler
    .long SPI1_IRQHandler
    .long I2S0_Tx_IRQHandler
    .long I2S0_Rx_IRQHandler
    .long 0
    .long UART0_RX_TX_IRQHandler
    .long UART0_ERR_IRQHandler
    .long UART1_RX_TX_IRQHandler
    .long UART1_ERR_IRQHandler
    .long UART2_RX_TX_IRQHandler
    .long UART2_ERR_IRQHandler
    .long UART3_RX_TX_IRQHandler
    .long UART3_ERR_IRQHandler
    .long ADC0_IRQHandler
    .long CMP0_IRQHandler
    .long CMP1_IRQHandler
    .long FTM0_IRQHandler
    .long FTM1_IRQHandler
    .long FTM2_IRQHandler
    .long CMT_IRQHandler
    .long RTC_IRQHandler
    .long RTC_Seconds_IRQHandler
    .long PIT0_IRQHandler
    .long PIT1_IRQHandler
    .long PIT2_IRQHandler
    .long PIT3_IRQHandler
    .long PDB0_IRQHandler
    .long USB0_IRQHandler
    .long USBDCD_IRQHandler
    .long 0
    .long DAC0_IRQHandler
    .long 0
    .long LPTMR0_IRQHandler
    .long PORTA_IRQHandler
    .long PORTB_IRQHandler
    .long PORTC_IRQHandler
    .long PORTD_IRQHandler
    .long PORTE_IRQHandler
    .long 0
    .long SPI2_IRQHandler
    .long UART4_RX_TX_IRQHandler
    .long UART4_ERR_IRQHandler
    .long 0
    .long 0
    .long CMP2_IRQHandler
    .long FTM3_IRQHandler
    .long DAC1_IRQHandler
    .long ADC1_IRQHandler
    .long I2C2_IRQHandler
    .long CAN0_ORed_Message_buffer_IRQHandler
    .long CAN0_Bus_Off_IRQHandler
    .long CAN0_Error_IRQHandler
    .long CAN0_Tx_Warning_IRQHandler
    .long CAN0_Rx_Warning_IRQHandler
    .long CAN0_Wake_Up_IRQHandler
    .long SDHC_IRQHandler
    .long ENET_1588_Timer_IRQHandler
    .long ENET_Transmit_IRQHandler
    .long ENET_Receive_IRQHandler
    .long ENET_Error_IRQHandler
    .long LPUART0_IRQHandler
    .long TSI0_IRQHandler
    .long TPM1_IRQHandler
    .long TPM2_IRQHandler
    .long USBHSDCD_IRQHandler
    .long I2C3_IRQHandler
    .long CMP3_IRQHandler
    .long USBHS_IRQHandler
    .long CAN1_ORed_Message_buffer_IRQHandler
    .long CAN1_Bus_Off_IRQHandler
    .long CAN1_Error_IRQHandler
    .long CAN1_Tx_Warning_IRQHandler
    .long CAN1_Rx_Warning_IRQHandler
    .long CAN1_Wake_Up_IRQHandler

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
    IRQ DMA0_DMA16_IRQHandler
    IRQ DMA1_DMA17_IRQHandler
    IRQ DMA2_DMA18_IRQHandler
    IRQ DMA3_DMA19_IRQHandler
    IRQ DMA4_DMA20_IRQHandler
    IRQ DMA5_DMA21_IRQHandler
    IRQ DMA6_DMA22_IRQHandler
    IRQ DMA7_DMA23_IRQHandler
    IRQ DMA8_DMA24_IRQHandler
    IRQ DMA9_DMA25_IRQHandler
    IRQ DMA10_DMA26_IRQHandler
    IRQ DMA11_DMA27_IRQHandler
    IRQ DMA12_DMA28_IRQHandler
    IRQ DMA13_DMA29_IRQHandler
    IRQ DMA14_DMA30_IRQHandler
    IRQ DMA15_DMA31_IRQHandler
    IRQ DMA_Error_IRQHandler
    IRQ MCM_IRQHandler
    IRQ FTFE_IRQHandler
    IRQ Read_Collision_IRQHandler
    IRQ LVD_LVW_IRQHandler
    IRQ LLWU_IRQHandler
    IRQ WDOG_EWM_IRQHandler
    IRQ RNG_IRQHandler
    IRQ I2C0_IRQHandler
    IRQ I2C1_IRQHandler
    IRQ SPI0_IRQHandler
    IRQ SPI1_IRQHandler
    IRQ I2S0_Tx_IRQHandler
    IRQ I2S0_Rx_IRQHandler
    IRQ UART0_RX_TX_IRQHandler
    IRQ UART0_ERR_IRQHandler
    IRQ UART1_RX_TX_IRQHandler
    IRQ UART1_ERR_IRQHandler
    IRQ UART2_RX_TX_IRQHandler
    IRQ UART2_ERR_IRQHandler
    IRQ UART3_RX_TX_IRQHandler
    IRQ UART3_ERR_IRQHandler
    IRQ ADC0_IRQHandler
    IRQ CMP0_IRQHandler
    IRQ CMP1_IRQHandler
    IRQ FTM0_IRQHandler
    IRQ FTM1_IRQHandler
    IRQ FTM2_IRQHandler
    IRQ CMT_IRQHandler
    IRQ RTC_IRQHandler
    IRQ RTC_Seconds_IRQHandler
    IRQ PIT0_IRQHandler
    IRQ PIT1_IRQHandler
    IRQ PIT2_IRQHandler
    IRQ PIT3_IRQHandler
    IRQ PDB0_IRQHandler
    IRQ USB0_IRQHandler
    IRQ USBDCD_IRQHandler
    IRQ DAC0_IRQHandler
    IRQ LPTMR0_IRQHandler
    IRQ PORTA_IRQHandler
    IRQ PORTB_IRQHandler
    IRQ PORTC_IRQHandler
    IRQ PORTD_IRQHandler
    IRQ PORTE_IRQHandler
    IRQ SPI2_IRQHandler
    IRQ UART4_RX_TX_IRQHandler
    IRQ UART4_ERR_IRQHandler
    IRQ CMP2_IRQHandler
    IRQ FTM3_IRQHandler
    IRQ DAC1_IRQHandler
    IRQ ADC1_IRQHandler
    IRQ I2C2_IRQHandler
    IRQ CAN0_ORed_Message_buffer_IRQHandler
    IRQ CAN0_Bus_Off_IRQHandler
    IRQ CAN0_Error_IRQHandler
    IRQ CAN0_Tx_Warning_IRQHandler
    IRQ CAN0_Rx_Warning_IRQHandler
    IRQ CAN0_Wake_Up_IRQHandler
    IRQ SDHC_IRQHandler
    IRQ ENET_1588_Timer_IRQHandler
    IRQ ENET_Transmit_IRQHandler
    IRQ ENET_Receive_IRQHandler
    IRQ ENET_Error_IRQHandler
    IRQ LPUART0_IRQHandler
    IRQ TSI0_IRQHandler
    IRQ TPM1_IRQHandler
    IRQ TPM2_IRQHandler
    IRQ USBHSDCD_IRQHandler
    IRQ I2C3_IRQHandler
    IRQ CMP3_IRQHandler
    IRQ USBHS_IRQHandler
    IRQ CAN1_ORed_Message_buffer_IRQHandler
    IRQ CAN1_Bus_Off_IRQHandler
    IRQ CAN1_Error_IRQHandler
    IRQ CAN1_Tx_Warning_IRQHandler
    IRQ CAN1_Rx_Warning_IRQHandler
    IRQ CAN1_Wake_Up_IRQHandler
