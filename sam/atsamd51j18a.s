// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from ATSAMD51J18A.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/Atmel

// Microchip ATSAMD51J18A Microcontroller
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
    .long PM_IRQHandler
    .long MCLK_IRQHandler
    .long OSCCTRL_XOSC0_IRQHandler
    .long OSCCTRL_XOSC1_IRQHandler
    .long OSCCTRL_DFLL_IRQHandler
    .long OSCCTRL_DPLL0_IRQHandler
    .long OSCCTRL_DPLL1_IRQHandler
    .long OSC32KCTRL_IRQHandler
    .long SUPC_OTHER_IRQHandler
    .long SUPC_BODDET_IRQHandler
    .long WDT_IRQHandler
    .long RTC_IRQHandler
    .long EIC_EXTINT_0_IRQHandler
    .long EIC_EXTINT_1_IRQHandler
    .long EIC_EXTINT_2_IRQHandler
    .long EIC_EXTINT_3_IRQHandler
    .long EIC_EXTINT_4_IRQHandler
    .long EIC_EXTINT_5_IRQHandler
    .long EIC_EXTINT_6_IRQHandler
    .long EIC_EXTINT_7_IRQHandler
    .long EIC_EXTINT_8_IRQHandler
    .long EIC_EXTINT_9_IRQHandler
    .long EIC_EXTINT_10_IRQHandler
    .long EIC_EXTINT_11_IRQHandler
    .long EIC_EXTINT_12_IRQHandler
    .long EIC_EXTINT_13_IRQHandler
    .long EIC_EXTINT_14_IRQHandler
    .long EIC_EXTINT_15_IRQHandler
    .long FREQM_IRQHandler
    .long NVMCTRL_0_IRQHandler
    .long NVMCTRL_1_IRQHandler
    .long DMAC_0_IRQHandler
    .long DMAC_1_IRQHandler
    .long DMAC_2_IRQHandler
    .long DMAC_3_IRQHandler
    .long DMAC_OTHER_IRQHandler
    .long EVSYS_0_IRQHandler
    .long EVSYS_1_IRQHandler
    .long EVSYS_2_IRQHandler
    .long EVSYS_3_IRQHandler
    .long EVSYS_OTHER_IRQHandler
    .long PAC_IRQHandler
    .long 0
    .long 0
    .long 0
    .long RAMECC_IRQHandler
    .long SERCOM0_0_IRQHandler
    .long SERCOM0_1_IRQHandler
    .long SERCOM0_2_IRQHandler
    .long SERCOM0_OTHER_IRQHandler
    .long SERCOM1_0_IRQHandler
    .long SERCOM1_1_IRQHandler
    .long SERCOM1_2_IRQHandler
    .long SERCOM1_OTHER_IRQHandler
    .long SERCOM2_0_IRQHandler
    .long SERCOM2_1_IRQHandler
    .long SERCOM2_2_IRQHandler
    .long SERCOM2_OTHER_IRQHandler
    .long SERCOM3_0_IRQHandler
    .long SERCOM3_1_IRQHandler
    .long SERCOM3_2_IRQHandler
    .long SERCOM3_OTHER_IRQHandler
    .long SERCOM4_0_IRQHandler
    .long SERCOM4_1_IRQHandler
    .long SERCOM4_2_IRQHandler
    .long SERCOM4_OTHER_IRQHandler
    .long SERCOM5_0_IRQHandler
    .long SERCOM5_1_IRQHandler
    .long SERCOM5_2_IRQHandler
    .long SERCOM5_OTHER_IRQHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long USB_OTHER_IRQHandler
    .long USB_SOF_HSOF_IRQHandler
    .long USB_TRCPT0_IRQHandler
    .long USB_TRCPT1_IRQHandler
    .long 0
    .long TCC0_OTHER_IRQHandler
    .long TCC0_MC0_IRQHandler
    .long TCC0_MC1_IRQHandler
    .long TCC0_MC2_IRQHandler
    .long TCC0_MC3_IRQHandler
    .long TCC0_MC4_IRQHandler
    .long TCC0_MC5_IRQHandler
    .long TCC1_OTHER_IRQHandler
    .long TCC1_MC0_IRQHandler
    .long TCC1_MC1_IRQHandler
    .long TCC1_MC2_IRQHandler
    .long TCC1_MC3_IRQHandler
    .long TCC2_OTHER_IRQHandler
    .long TCC2_MC0_IRQHandler
    .long TCC2_MC1_IRQHandler
    .long TCC2_MC2_IRQHandler
    .long TCC3_OTHER_IRQHandler
    .long TCC3_MC0_IRQHandler
    .long TCC3_MC1_IRQHandler
    .long TCC4_OTHER_IRQHandler
    .long TCC4_MC0_IRQHandler
    .long TCC4_MC1_IRQHandler
    .long TC0_IRQHandler
    .long TC1_IRQHandler
    .long TC2_IRQHandler
    .long TC3_IRQHandler
    .long TC4_IRQHandler
    .long TC5_IRQHandler
    .long 0
    .long 0
    .long PDEC_OTHER_IRQHandler
    .long PDEC_MC0_IRQHandler
    .long PDEC_MC1_IRQHandler
    .long ADC0_OTHER_IRQHandler
    .long ADC0_RESRDY_IRQHandler
    .long ADC1_OTHER_IRQHandler
    .long ADC1_RESRDY_IRQHandler
    .long AC_IRQHandler
    .long DAC_OTHER_IRQHandler
    .long DAC_EMPTY_0_IRQHandler
    .long DAC_EMPTY_1_IRQHandler
    .long DAC_RESRDY_0_IRQHandler
    .long DAC_RESRDY_1_IRQHandler
    .long I2S_IRQHandler
    .long PCC_IRQHandler
    .long AES_IRQHandler
    .long TRNG_IRQHandler
    .long ICM_IRQHandler
    .long 0
    .long QSPI_IRQHandler
    .long SDHC0_IRQHandler

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
    IRQ PM_IRQHandler
    IRQ MCLK_IRQHandler
    IRQ OSCCTRL_XOSC0_IRQHandler
    IRQ OSCCTRL_XOSC1_IRQHandler
    IRQ OSCCTRL_DFLL_IRQHandler
    IRQ OSCCTRL_DPLL0_IRQHandler
    IRQ OSCCTRL_DPLL1_IRQHandler
    IRQ OSC32KCTRL_IRQHandler
    IRQ SUPC_OTHER_IRQHandler
    IRQ SUPC_BODDET_IRQHandler
    IRQ WDT_IRQHandler
    IRQ RTC_IRQHandler
    IRQ EIC_EXTINT_0_IRQHandler
    IRQ EIC_EXTINT_1_IRQHandler
    IRQ EIC_EXTINT_2_IRQHandler
    IRQ EIC_EXTINT_3_IRQHandler
    IRQ EIC_EXTINT_4_IRQHandler
    IRQ EIC_EXTINT_5_IRQHandler
    IRQ EIC_EXTINT_6_IRQHandler
    IRQ EIC_EXTINT_7_IRQHandler
    IRQ EIC_EXTINT_8_IRQHandler
    IRQ EIC_EXTINT_9_IRQHandler
    IRQ EIC_EXTINT_10_IRQHandler
    IRQ EIC_EXTINT_11_IRQHandler
    IRQ EIC_EXTINT_12_IRQHandler
    IRQ EIC_EXTINT_13_IRQHandler
    IRQ EIC_EXTINT_14_IRQHandler
    IRQ EIC_EXTINT_15_IRQHandler
    IRQ FREQM_IRQHandler
    IRQ NVMCTRL_0_IRQHandler
    IRQ NVMCTRL_1_IRQHandler
    IRQ DMAC_0_IRQHandler
    IRQ DMAC_1_IRQHandler
    IRQ DMAC_2_IRQHandler
    IRQ DMAC_3_IRQHandler
    IRQ DMAC_OTHER_IRQHandler
    IRQ EVSYS_0_IRQHandler
    IRQ EVSYS_1_IRQHandler
    IRQ EVSYS_2_IRQHandler
    IRQ EVSYS_3_IRQHandler
    IRQ EVSYS_OTHER_IRQHandler
    IRQ PAC_IRQHandler
    IRQ RAMECC_IRQHandler
    IRQ SERCOM0_0_IRQHandler
    IRQ SERCOM0_1_IRQHandler
    IRQ SERCOM0_2_IRQHandler
    IRQ SERCOM0_OTHER_IRQHandler
    IRQ SERCOM1_0_IRQHandler
    IRQ SERCOM1_1_IRQHandler
    IRQ SERCOM1_2_IRQHandler
    IRQ SERCOM1_OTHER_IRQHandler
    IRQ SERCOM2_0_IRQHandler
    IRQ SERCOM2_1_IRQHandler
    IRQ SERCOM2_2_IRQHandler
    IRQ SERCOM2_OTHER_IRQHandler
    IRQ SERCOM3_0_IRQHandler
    IRQ SERCOM3_1_IRQHandler
    IRQ SERCOM3_2_IRQHandler
    IRQ SERCOM3_OTHER_IRQHandler
    IRQ SERCOM4_0_IRQHandler
    IRQ SERCOM4_1_IRQHandler
    IRQ SERCOM4_2_IRQHandler
    IRQ SERCOM4_OTHER_IRQHandler
    IRQ SERCOM5_0_IRQHandler
    IRQ SERCOM5_1_IRQHandler
    IRQ SERCOM5_2_IRQHandler
    IRQ SERCOM5_OTHER_IRQHandler
    IRQ USB_OTHER_IRQHandler
    IRQ USB_SOF_HSOF_IRQHandler
    IRQ USB_TRCPT0_IRQHandler
    IRQ USB_TRCPT1_IRQHandler
    IRQ TCC0_OTHER_IRQHandler
    IRQ TCC0_MC0_IRQHandler
    IRQ TCC0_MC1_IRQHandler
    IRQ TCC0_MC2_IRQHandler
    IRQ TCC0_MC3_IRQHandler
    IRQ TCC0_MC4_IRQHandler
    IRQ TCC0_MC5_IRQHandler
    IRQ TCC1_OTHER_IRQHandler
    IRQ TCC1_MC0_IRQHandler
    IRQ TCC1_MC1_IRQHandler
    IRQ TCC1_MC2_IRQHandler
    IRQ TCC1_MC3_IRQHandler
    IRQ TCC2_OTHER_IRQHandler
    IRQ TCC2_MC0_IRQHandler
    IRQ TCC2_MC1_IRQHandler
    IRQ TCC2_MC2_IRQHandler
    IRQ TCC3_OTHER_IRQHandler
    IRQ TCC3_MC0_IRQHandler
    IRQ TCC3_MC1_IRQHandler
    IRQ TCC4_OTHER_IRQHandler
    IRQ TCC4_MC0_IRQHandler
    IRQ TCC4_MC1_IRQHandler
    IRQ TC0_IRQHandler
    IRQ TC1_IRQHandler
    IRQ TC2_IRQHandler
    IRQ TC3_IRQHandler
    IRQ TC4_IRQHandler
    IRQ TC5_IRQHandler
    IRQ PDEC_OTHER_IRQHandler
    IRQ PDEC_MC0_IRQHandler
    IRQ PDEC_MC1_IRQHandler
    IRQ ADC0_OTHER_IRQHandler
    IRQ ADC0_RESRDY_IRQHandler
    IRQ ADC1_OTHER_IRQHandler
    IRQ ADC1_RESRDY_IRQHandler
    IRQ AC_IRQHandler
    IRQ DAC_OTHER_IRQHandler
    IRQ DAC_EMPTY_0_IRQHandler
    IRQ DAC_EMPTY_1_IRQHandler
    IRQ DAC_RESRDY_0_IRQHandler
    IRQ DAC_RESRDY_1_IRQHandler
    IRQ I2S_IRQHandler
    IRQ PCC_IRQHandler
    IRQ AES_IRQHandler
    IRQ TRNG_IRQHandler
    IRQ ICM_IRQHandler
    IRQ QSPI_IRQHandler
    IRQ SDHC0_IRQHandler
