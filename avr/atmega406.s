; Automatically generated file. DO NOT EDIT.
; Generated by gen-device-avr.go from ATmega406.atdf, see http://packs.download.atmel.com/

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
    jmp __vector_BPINT
    jmp __vector_INT0
    jmp __vector_INT1
    jmp __vector_INT2
    jmp __vector_INT3
    jmp __vector_PCINT0
    jmp __vector_PCINT1
    jmp __vector_WDT
    jmp __vector_WAKE_UP
    jmp __vector_TIM1_COMP
    jmp __vector_TIM1_OVF
    jmp __vector_TIM0_COMPA
    jmp __vector_TIM0_COMPB
    jmp __vector_TIM0_OVF
    jmp __vector_TWI_BUS_CD
    jmp __vector_TWI
    jmp __vector_VADC
    jmp __vector_CCADC_CONV
    jmp __vector_CCADC_REG_CUR
    jmp __vector_CCADC_ACC
    jmp __vector_EE_READY
    jmp __vector_SPM_READY

    ; Define default implementations for interrupts, redirecting to
    ; __vector_default when not implemented.
    IRQ __vector_RESET
    IRQ __vector_BPINT
    IRQ __vector_INT0
    IRQ __vector_INT1
    IRQ __vector_INT2
    IRQ __vector_INT3
    IRQ __vector_PCINT0
    IRQ __vector_PCINT1
    IRQ __vector_WDT
    IRQ __vector_WAKE_UP
    IRQ __vector_TIM1_COMP
    IRQ __vector_TIM1_OVF
    IRQ __vector_TIM0_COMPA
    IRQ __vector_TIM0_COMPB
    IRQ __vector_TIM0_OVF
    IRQ __vector_TWI_BUS_CD
    IRQ __vector_TWI
    IRQ __vector_VADC
    IRQ __vector_CCADC_CONV
    IRQ __vector_CCADC_REG_CUR
    IRQ __vector_CCADC_ACC
    IRQ __vector_EE_READY
    IRQ __vector_SPM_READY
