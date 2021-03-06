; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATtiny1634.atdf, see http://packs.download.atmel.com/

; This is the default handler for interrupts, if triggered but not defined.
; Sleep inside so that an accidentally triggered interrupt won't drain the
; battery of a battery-powered device.
.section .text.__vector_default
.global  __vector_default
__vector_default:
    sleep
    rjmp __vector_default

; Avoid the need for repeated .weak and .set instructions.
.macro IRQ handler
    .weak  \handler
    .set   \handler, __vector_default
.endm

; The interrupt vector of this device. Must be placed at address 0 by the linker.
.section .vectors
.global  __vectors
    jmp __vector_RESET
    jmp __vector_INT0
    jmp __vector_PCINT0
    jmp __vector_PCINT1
    jmp __vector_PCINT2
    jmp __vector_WDT
    jmp __vector_TIMER1_CAPT
    jmp __vector_TIMER1_COMPA
    jmp __vector_TIMER1_COMPB
    jmp __vector_TIMER1_OVF
    jmp __vector_TIMER0_COMPA
    jmp __vector_TIMER0_COMPB
    jmp __vector_TIMER0_OVF
    jmp __vector_ANA_COMP
    jmp __vector_ADC
    jmp __vector_USART0_START
    jmp __vector_USART0_RX
    jmp __vector_USART0_UDRE
    jmp __vector_USART0_TX
    jmp __vector_USART1_START
    jmp __vector_USART1_RX
    jmp __vector_USART1_UDRE
    jmp __vector_USART1_TX
    jmp __vector_USI_START
    jmp __vector_USI_OVERFLOW
    jmp __vector_TWI_SLAVE
    jmp __vector_EE_RDY
    jmp __vector_QTRIP

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_RESET
    IRQ __vector_INT0
    IRQ __vector_PCINT0
    IRQ __vector_PCINT1
    IRQ __vector_PCINT2
    IRQ __vector_WDT
    IRQ __vector_TIMER1_CAPT
    IRQ __vector_TIM1_CAPT
    IRQ __vector_TIMER1_COMPA
    IRQ __vector_TIM1_COMPA
    IRQ __vector_TIMER1_COMPB
    IRQ __vector_TIM1_COMPB
    IRQ __vector_TIMER1_OVF
    IRQ __vector_TIM1_OVF
    IRQ __vector_TIMER0_COMPA
    IRQ __vector_TIM0_COMPA
    IRQ __vector_TIMER0_COMPB
    IRQ __vector_TIM0_COMPB
    IRQ __vector_TIMER0_OVF
    IRQ __vector_TIM0_OVF
    IRQ __vector_ANA_COMP
    IRQ __vector_ADC
    IRQ __vector_ADC_READY
    IRQ __vector_USART0_START
    IRQ __vector_USART0_RXS
    IRQ __vector_USART0_RX
    IRQ __vector_USART0_RXC
    IRQ __vector_USART0_UDRE
    IRQ __vector_USART0_DRE
    IRQ __vector_USART0_TX
    IRQ __vector_USART0_TXC
    IRQ __vector_USART1_START
    IRQ __vector_USART1_RXS
    IRQ __vector_USART1_RX
    IRQ __vector_USART1_RXC
    IRQ __vector_USART1_UDRE
    IRQ __vector_USART1_DRE
    IRQ __vector_USART1_TX
    IRQ __vector_USART1_TXC
    IRQ __vector_USI_START
    IRQ __vector_USI_STR
    IRQ __vector_USI_OVERFLOW
    IRQ __vector_USI_OVF
    IRQ __vector_TWI_SLAVE
    IRQ __vector_TWI
    IRQ __vector_EE_RDY
    IRQ __vector_QTRIP
