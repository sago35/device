// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-svd.go from MK22F51212.svd, see https://github.com/posborne/cmsis-svd/tree/master/data/NXP

// MK22F51212 NXP Microcontroller
//
//     Redistribution and use in source and binary forms, with or without modification,\nare permitted provided that the following conditions are met:
//     o Redistributions of source code must retain the above copyright notice, this list
//     of conditions and the following disclaimer.
//     o Redistributions in binary form must reproduce the above copyright notice, this
//     list of conditions and the following disclaimer in the documentation and/or
//     other materials provided with the distribution.
//     o Neither the name of the copyright holder nor the names of its
//     contributors may be used to endorse or promote products derived from this
//     software without specific prior written permission.
//     THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
//     ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
//     WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//     DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
//     ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
//     (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
//     LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
//     ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
//     (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
//     SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
    .long DMA0_IRQHandler
    .long DMA1_IRQHandler
    .long DMA2_IRQHandler
    .long DMA3_IRQHandler
    .long DMA4_IRQHandler
    .long DMA5_IRQHandler
    .long DMA6_IRQHandler
    .long DMA7_IRQHandler
    .long DMA8_IRQHandler
    .long DMA9_IRQHandler
    .long DMA10_IRQHandler
    .long DMA11_IRQHandler
    .long DMA12_IRQHandler
    .long DMA13_IRQHandler
    .long DMA14_IRQHandler
    .long DMA15_IRQHandler
    .long DMA_Error_IRQHandler
    .long MCM_IRQHandler
    .long FTF_IRQHandler
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
    .long LPUART0_IRQHandler
    .long UART0_RX_TX_IRQHandler
    .long UART0_ERR_IRQHandler
    .long UART1_RX_TX_IRQHandler
    .long UART1_ERR_IRQHandler
    .long UART2_RX_TX_IRQHandler
    .long UART2_ERR_IRQHandler
    .long 0
    .long 0
    .long ADC0_IRQHandler
    .long CMP0_IRQHandler
    .long CMP1_IRQHandler
    .long FTM0_IRQHandler
    .long FTM1_IRQHandler
    .long FTM2_IRQHandler
    .long 0
    .long RTC_IRQHandler
    .long RTC_Seconds_IRQHandler
    .long PIT0_IRQHandler
    .long PIT1_IRQHandler
    .long PIT2_IRQHandler
    .long PIT3_IRQHandler
    .long PDB0_IRQHandler
    .long USB0_IRQHandler
    .long 0
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
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long 0
    .long FTM3_IRQHandler
    .long DAC1_IRQHandler
    .long ADC1_IRQHandler

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
    IRQ DMA0_IRQHandler
    IRQ DMA1_IRQHandler
    IRQ DMA2_IRQHandler
    IRQ DMA3_IRQHandler
    IRQ DMA4_IRQHandler
    IRQ DMA5_IRQHandler
    IRQ DMA6_IRQHandler
    IRQ DMA7_IRQHandler
    IRQ DMA8_IRQHandler
    IRQ DMA9_IRQHandler
    IRQ DMA10_IRQHandler
    IRQ DMA11_IRQHandler
    IRQ DMA12_IRQHandler
    IRQ DMA13_IRQHandler
    IRQ DMA14_IRQHandler
    IRQ DMA15_IRQHandler
    IRQ DMA_Error_IRQHandler
    IRQ MCM_IRQHandler
    IRQ FTF_IRQHandler
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
    IRQ LPUART0_IRQHandler
    IRQ UART0_RX_TX_IRQHandler
    IRQ UART0_ERR_IRQHandler
    IRQ UART1_RX_TX_IRQHandler
    IRQ UART1_ERR_IRQHandler
    IRQ UART2_RX_TX_IRQHandler
    IRQ UART2_ERR_IRQHandler
    IRQ ADC0_IRQHandler
    IRQ CMP0_IRQHandler
    IRQ CMP1_IRQHandler
    IRQ FTM0_IRQHandler
    IRQ FTM1_IRQHandler
    IRQ FTM2_IRQHandler
    IRQ RTC_IRQHandler
    IRQ RTC_Seconds_IRQHandler
    IRQ PIT0_IRQHandler
    IRQ PIT1_IRQHandler
    IRQ PIT2_IRQHandler
    IRQ PIT3_IRQHandler
    IRQ PDB0_IRQHandler
    IRQ USB0_IRQHandler
    IRQ DAC0_IRQHandler
    IRQ LPTMR0_IRQHandler
    IRQ PORTA_IRQHandler
    IRQ PORTB_IRQHandler
    IRQ PORTC_IRQHandler
    IRQ PORTD_IRQHandler
    IRQ PORTE_IRQHandler
    IRQ FTM3_IRQHandler
    IRQ DAC1_IRQHandler
    IRQ ADC1_IRQHandler
