// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from MIMXRT1062.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// MIMXRT1062DVL6A
//
//     Copyright 2016-2019 NXP All rights reserved. SPDX-License-Identifier: BSD-3-Clause

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
    .long DMA_ERROR_IRQHandler
    .long CTI0_ERROR_IRQHandler
    .long CTI1_ERROR_IRQHandler
    .long CORE_IRQHandler
    .long LPUART1_IRQHandler
    .long LPUART2_IRQHandler
    .long LPUART3_IRQHandler
    .long LPUART4_IRQHandler
    .long LPUART5_IRQHandler
    .long LPUART6_IRQHandler
    .long LPUART7_IRQHandler
    .long LPUART8_IRQHandler
    .long LPI2C1_IRQHandler
    .long LPI2C2_IRQHandler
    .long LPI2C3_IRQHandler
    .long LPI2C4_IRQHandler
    .long LPSPI1_IRQHandler
    .long LPSPI2_IRQHandler
    .long LPSPI3_IRQHandler
    .long LPSPI4_IRQHandler
    .long CAN1_IRQHandler
    .long CAN2_IRQHandler
    .long FLEXRAM_IRQHandler
    .long KPP_IRQHandler
    .long TSC_DIG_IRQHandler
    .long GPR_IRQ_IRQHandler
    .long LCDIF_IRQHandler
    .long CSI_IRQHandler
    .long PXP_IRQHandler
    .long WDOG2_IRQHandler
    .long SNVS_HP_WRAPPER_IRQHandler
    .long SNVS_HP_WRAPPER_TZ_IRQHandler
    .long SNVS_LP_WRAPPER_IRQHandler
    .long CSU_IRQHandler
    .long DCP_IRQHandler
    .long DCP_VMI_IRQHandler
    .long Reserved68_IRQHandler
    .long TRNG_IRQHandler
    .long SJC_IRQHandler
    .long BEE_IRQHandler
    .long SAI1_IRQHandler
    .long SAI2_IRQHandler
    .long SAI3_RX_IRQHandler
    .long SAI3_TX_IRQHandler
    .long SPDIF_IRQHandler
    .long PMU_EVENT_IRQHandler
    .long Reserved78_IRQHandler
    .long TEMP_LOW_HIGH_IRQHandler
    .long TEMP_PANIC_IRQHandler
    .long USB_PHY1_IRQHandler
    .long USB_PHY2_IRQHandler
    .long ADC1_IRQHandler
    .long ADC2_IRQHandler
    .long DCDC_IRQHandler
    .long Reserved86_IRQHandler
    .long Reserved87_IRQHandler
    .long GPIO1_INT0_IRQHandler
    .long GPIO1_INT1_IRQHandler
    .long GPIO1_INT2_IRQHandler
    .long GPIO1_INT3_IRQHandler
    .long GPIO1_INT4_IRQHandler
    .long GPIO1_INT5_IRQHandler
    .long GPIO1_INT6_IRQHandler
    .long GPIO1_INT7_IRQHandler
    .long GPIO1_Combined_0_15_IRQHandler
    .long GPIO1_Combined_16_31_IRQHandler
    .long GPIO2_Combined_0_15_IRQHandler
    .long GPIO2_Combined_16_31_IRQHandler
    .long GPIO3_Combined_0_15_IRQHandler
    .long GPIO3_Combined_16_31_IRQHandler
    .long GPIO4_Combined_0_15_IRQHandler
    .long GPIO4_Combined_16_31_IRQHandler
    .long GPIO5_Combined_0_15_IRQHandler
    .long GPIO5_Combined_16_31_IRQHandler
    .long FLEXIO1_IRQHandler
    .long FLEXIO2_IRQHandler
    .long WDOG1_IRQHandler
    .long RTWDOG_IRQHandler
    .long EWM_IRQHandler
    .long CCM_1_IRQHandler
    .long CCM_2_IRQHandler
    .long GPC_IRQHandler
    .long SRC_IRQHandler
    .long Reserved115_IRQHandler
    .long GPT1_IRQHandler
    .long GPT2_IRQHandler
    .long PWM1_0_IRQHandler
    .long PWM1_1_IRQHandler
    .long PWM1_2_IRQHandler
    .long PWM1_3_IRQHandler
    .long PWM1_FAULT_IRQHandler
    .long FLEXSPI2_IRQHandler
    .long FLEXSPI_IRQHandler
    .long SEMC_IRQHandler
    .long USDHC1_IRQHandler
    .long USDHC2_IRQHandler
    .long USB_OTG2_IRQHandler
    .long USB_OTG1_IRQHandler
    .long ENET_IRQHandler
    .long ENET_1588_Timer_IRQHandler
    .long XBAR1_IRQ_0_1_IRQHandler
    .long XBAR1_IRQ_2_3_IRQHandler
    .long ADC_ETC_IRQ0_IRQHandler
    .long ADC_ETC_IRQ1_IRQHandler
    .long ADC_ETC_IRQ2_IRQHandler
    .long ADC_ETC_ERROR_IRQ_IRQHandler
    .long PIT_IRQHandler
    .long ACMP1_IRQHandler
    .long ACMP2_IRQHandler
    .long ACMP3_IRQHandler
    .long ACMP4_IRQHandler
    .long Reserved143_IRQHandler
    .long Reserved144_IRQHandler
    .long ENC1_IRQHandler
    .long ENC2_IRQHandler
    .long ENC3_IRQHandler
    .long ENC4_IRQHandler
    .long TMR1_IRQHandler
    .long TMR2_IRQHandler
    .long TMR3_IRQHandler
    .long TMR4_IRQHandler
    .long PWM2_0_IRQHandler
    .long PWM2_1_IRQHandler
    .long PWM2_2_IRQHandler
    .long PWM2_3_IRQHandler
    .long PWM2_FAULT_IRQHandler
    .long PWM3_0_IRQHandler
    .long PWM3_1_IRQHandler
    .long PWM3_2_IRQHandler
    .long PWM3_3_IRQHandler
    .long PWM3_FAULT_IRQHandler
    .long PWM4_0_IRQHandler
    .long PWM4_1_IRQHandler
    .long PWM4_2_IRQHandler
    .long PWM4_3_IRQHandler
    .long PWM4_FAULT_IRQHandler
    .long ENET2_IRQHandler
    .long ENET2_1588_Timer_IRQHandler
    .long CAN3_IRQHandler
    .long Reserved171_IRQHandler
    .long FLEXIO3_IRQHandler
    .long GPIO6_7_8_9_IRQHandler

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
    IRQ DMA_ERROR_IRQHandler
    IRQ CTI0_ERROR_IRQHandler
    IRQ CTI1_ERROR_IRQHandler
    IRQ CORE_IRQHandler
    IRQ LPUART1_IRQHandler
    IRQ LPUART2_IRQHandler
    IRQ LPUART3_IRQHandler
    IRQ LPUART4_IRQHandler
    IRQ LPUART5_IRQHandler
    IRQ LPUART6_IRQHandler
    IRQ LPUART7_IRQHandler
    IRQ LPUART8_IRQHandler
    IRQ LPI2C1_IRQHandler
    IRQ LPI2C2_IRQHandler
    IRQ LPI2C3_IRQHandler
    IRQ LPI2C4_IRQHandler
    IRQ LPSPI1_IRQHandler
    IRQ LPSPI2_IRQHandler
    IRQ LPSPI3_IRQHandler
    IRQ LPSPI4_IRQHandler
    IRQ CAN1_IRQHandler
    IRQ CAN2_IRQHandler
    IRQ FLEXRAM_IRQHandler
    IRQ KPP_IRQHandler
    IRQ TSC_DIG_IRQHandler
    IRQ GPR_IRQ_IRQHandler
    IRQ LCDIF_IRQHandler
    IRQ CSI_IRQHandler
    IRQ PXP_IRQHandler
    IRQ WDOG2_IRQHandler
    IRQ SNVS_HP_WRAPPER_IRQHandler
    IRQ SNVS_HP_WRAPPER_TZ_IRQHandler
    IRQ SNVS_LP_WRAPPER_IRQHandler
    IRQ CSU_IRQHandler
    IRQ DCP_IRQHandler
    IRQ DCP_VMI_IRQHandler
    IRQ Reserved68_IRQHandler
    IRQ TRNG_IRQHandler
    IRQ SJC_IRQHandler
    IRQ BEE_IRQHandler
    IRQ SAI1_IRQHandler
    IRQ SAI2_IRQHandler
    IRQ SAI3_RX_IRQHandler
    IRQ SAI3_TX_IRQHandler
    IRQ SPDIF_IRQHandler
    IRQ PMU_EVENT_IRQHandler
    IRQ Reserved78_IRQHandler
    IRQ TEMP_LOW_HIGH_IRQHandler
    IRQ TEMP_PANIC_IRQHandler
    IRQ USB_PHY1_IRQHandler
    IRQ USB_PHY2_IRQHandler
    IRQ ADC1_IRQHandler
    IRQ ADC2_IRQHandler
    IRQ DCDC_IRQHandler
    IRQ Reserved86_IRQHandler
    IRQ Reserved87_IRQHandler
    IRQ GPIO1_INT0_IRQHandler
    IRQ GPIO1_INT1_IRQHandler
    IRQ GPIO1_INT2_IRQHandler
    IRQ GPIO1_INT3_IRQHandler
    IRQ GPIO1_INT4_IRQHandler
    IRQ GPIO1_INT5_IRQHandler
    IRQ GPIO1_INT6_IRQHandler
    IRQ GPIO1_INT7_IRQHandler
    IRQ GPIO1_Combined_0_15_IRQHandler
    IRQ GPIO1_Combined_16_31_IRQHandler
    IRQ GPIO2_Combined_0_15_IRQHandler
    IRQ GPIO2_Combined_16_31_IRQHandler
    IRQ GPIO3_Combined_0_15_IRQHandler
    IRQ GPIO3_Combined_16_31_IRQHandler
    IRQ GPIO4_Combined_0_15_IRQHandler
    IRQ GPIO4_Combined_16_31_IRQHandler
    IRQ GPIO5_Combined_0_15_IRQHandler
    IRQ GPIO5_Combined_16_31_IRQHandler
    IRQ FLEXIO1_IRQHandler
    IRQ FLEXIO2_IRQHandler
    IRQ WDOG1_IRQHandler
    IRQ RTWDOG_IRQHandler
    IRQ EWM_IRQHandler
    IRQ CCM_1_IRQHandler
    IRQ CCM_2_IRQHandler
    IRQ GPC_IRQHandler
    IRQ SRC_IRQHandler
    IRQ Reserved115_IRQHandler
    IRQ GPT1_IRQHandler
    IRQ GPT2_IRQHandler
    IRQ PWM1_0_IRQHandler
    IRQ PWM1_1_IRQHandler
    IRQ PWM1_2_IRQHandler
    IRQ PWM1_3_IRQHandler
    IRQ PWM1_FAULT_IRQHandler
    IRQ FLEXSPI2_IRQHandler
    IRQ FLEXSPI_IRQHandler
    IRQ SEMC_IRQHandler
    IRQ USDHC1_IRQHandler
    IRQ USDHC2_IRQHandler
    IRQ USB_OTG2_IRQHandler
    IRQ USB_OTG1_IRQHandler
    IRQ ENET_IRQHandler
    IRQ ENET_1588_Timer_IRQHandler
    IRQ XBAR1_IRQ_0_1_IRQHandler
    IRQ XBAR1_IRQ_2_3_IRQHandler
    IRQ ADC_ETC_IRQ0_IRQHandler
    IRQ ADC_ETC_IRQ1_IRQHandler
    IRQ ADC_ETC_IRQ2_IRQHandler
    IRQ ADC_ETC_ERROR_IRQ_IRQHandler
    IRQ PIT_IRQHandler
    IRQ ACMP1_IRQHandler
    IRQ ACMP2_IRQHandler
    IRQ ACMP3_IRQHandler
    IRQ ACMP4_IRQHandler
    IRQ Reserved143_IRQHandler
    IRQ Reserved144_IRQHandler
    IRQ ENC1_IRQHandler
    IRQ ENC2_IRQHandler
    IRQ ENC3_IRQHandler
    IRQ ENC4_IRQHandler
    IRQ TMR1_IRQHandler
    IRQ TMR2_IRQHandler
    IRQ TMR3_IRQHandler
    IRQ TMR4_IRQHandler
    IRQ PWM2_0_IRQHandler
    IRQ PWM2_1_IRQHandler
    IRQ PWM2_2_IRQHandler
    IRQ PWM2_3_IRQHandler
    IRQ PWM2_FAULT_IRQHandler
    IRQ PWM3_0_IRQHandler
    IRQ PWM3_1_IRQHandler
    IRQ PWM3_2_IRQHandler
    IRQ PWM3_3_IRQHandler
    IRQ PWM3_FAULT_IRQHandler
    IRQ PWM4_0_IRQHandler
    IRQ PWM4_1_IRQHandler
    IRQ PWM4_2_IRQHandler
    IRQ PWM4_3_IRQHandler
    IRQ PWM4_FAULT_IRQHandler
    IRQ ENET2_IRQHandler
    IRQ ENET2_1588_Timer_IRQHandler
    IRQ CAN3_IRQHandler
    IRQ Reserved171_IRQHandler
    IRQ FLEXIO3_IRQHandler
    IRQ GPIO6_7_8_9_IRQHandler
