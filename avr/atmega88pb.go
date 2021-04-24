// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATmega88PB.atdf, see http://packs.download.atmel.com/

// +build avr,atmega88pb

// Device information for the ATmega88PB.
package avr

import (
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATmega88PB"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset and Watchdog Reset
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_PCINT0       = 3  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 4  // Pin Change Interrupt Request 0
	IRQ_PCINT2       = 5  // Pin Change Interrupt Request 1
	IRQ_WDT          = 6  // Watchdog Time-out Interrupt
	IRQ_TIMER2_COMPA = 7  // Timer/Counter2 Compare Match A
	IRQ_TIMER2_COMPB = 8  // Timer/Counter2 Compare Match A
	IRQ_TIMER2_OVF   = 9  // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 10 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 11 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 12 // Timer/Counter1 Compare Match B
	IRQ_TIMER1_OVF   = 13 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 14 // TimerCounter0 Compare Match A
	IRQ_TIMER0_COMPB = 15 // TimerCounter0 Compare Match B
	IRQ_TIMER0_OVF   = 16 // Timer/Couner0 Overflow
	IRQ_SPI_STC      = 17 // SPI Serial Transfer Complete
	IRQ_USART_RX     = 18 // USART Rx Complete
	IRQ_USART_UDRE   = 19 // USART, Data Register Empty
	IRQ_USART_TX     = 20 // USART Tx Complete
	IRQ_ADC          = 21 // ADC Conversion Complete
	IRQ_EE_READY     = 22 // EEPROM Ready
	IRQ_ANALOG_COMP  = 23 // Analog Comparator
	IRQ_TWI          = 24 // Two-wire Serial Interface
	IRQ_SPM_Ready    = 25 // Store Program Memory Read
	IRQ_USART_START  = 26 // USART Start Edge Interrupt
	IRQ_max          = 26 // Highest interrupt number on this device.
)

// Map interrupt numbers to function names.
// These aren't real calls, they're removed by the compiler.
var (
	_ = interrupt.Register(IRQ_RESET, "__vector_RESET")
	_ = interrupt.Register(IRQ_INT0, "__vector_INT0")
	_ = interrupt.Register(IRQ_INT1, "__vector_INT1")
	_ = interrupt.Register(IRQ_PCINT0, "__vector_PCINT0")
	_ = interrupt.Register(IRQ_PCINT1, "__vector_PCINT1")
	_ = interrupt.Register(IRQ_PCINT2, "__vector_PCINT2")
	_ = interrupt.Register(IRQ_WDT, "__vector_WDT")
	_ = interrupt.Register(IRQ_TIMER2_COMPA, "__vector_TIMER2_COMPA")
	_ = interrupt.Register(IRQ_TIMER2_COMPB, "__vector_TIMER2_COMPB")
	_ = interrupt.Register(IRQ_TIMER2_OVF, "__vector_TIMER2_OVF")
	_ = interrupt.Register(IRQ_TIMER1_CAPT, "__vector_TIMER1_CAPT")
	_ = interrupt.Register(IRQ_TIMER1_COMPA, "__vector_TIMER1_COMPA")
	_ = interrupt.Register(IRQ_TIMER1_COMPB, "__vector_TIMER1_COMPB")
	_ = interrupt.Register(IRQ_TIMER1_OVF, "__vector_TIMER1_OVF")
	_ = interrupt.Register(IRQ_TIMER0_COMPA, "__vector_TIMER0_COMPA")
	_ = interrupt.Register(IRQ_TIMER0_COMPB, "__vector_TIMER0_COMPB")
	_ = interrupt.Register(IRQ_TIMER0_OVF, "__vector_TIMER0_OVF")
	_ = interrupt.Register(IRQ_SPI_STC, "__vector_SPI_STC")
	_ = interrupt.Register(IRQ_USART_RX, "__vector_USART_RX")
	_ = interrupt.Register(IRQ_USART_UDRE, "__vector_USART_UDRE")
	_ = interrupt.Register(IRQ_USART_TX, "__vector_USART_TX")
	_ = interrupt.Register(IRQ_ADC, "__vector_ADC")
	_ = interrupt.Register(IRQ_EE_READY, "__vector_EE_READY")
	_ = interrupt.Register(IRQ_ANALOG_COMP, "__vector_ANALOG_COMP")
	_ = interrupt.Register(IRQ_TWI, "__vector_TWI")
	_ = interrupt.Register(IRQ_SPM_Ready, "__vector_SPM_Ready")
	_ = interrupt.Register(IRQ_USART_START, "__vector_USART_START")
)

// Peripherals.
var (
	// Fuses
	EXTENDED = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))
	HIGH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	LOW      = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Device ID
	DEVID0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf0)))
	DEVID1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf1)))
	DEVID2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf2)))
	DEVID3 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf3)))
	DEVID4 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf4)))
	DEVID5 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf5)))
	DEVID6 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf6)))
	DEVID7 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf7)))
	DEVID8 = (*volatile.Register8)(unsafe.Pointer(uintptr(0xf8)))

	// USART
	UDR0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc6)))
	UCSR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc0)))
	UCSR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc1)))
	UCSR0C = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc2)))
	UCSR0D = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc3)))
	UBRR0L = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc4)))
	UBRR0H = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc5)))

	// Two Wire Serial Interface
	TWAMR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbd)))
	TWBR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb8)))
	TWCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbc)))
	TWSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb9)))
	TWDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xbb)))
	TWAR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xba)))

	// Timer/Counter, 16-bit
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6f)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x82)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x88)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x89)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8a)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x8b)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x86)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x87)))

	// Timer/Counter, 8-bit Async
	TIMSK2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x70)))
	TIFR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	TCCR2A = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb0)))
	TCCR2B = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb1)))
	TCNT2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb2)))
	OCR2B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb4)))
	OCR2A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb3)))
	ASSR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb6)))

	// Analog-to-Digital Converter
	ADMUX  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7c)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x78)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x79)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7a)))
	ADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7b)))
	DIDR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7e)))

	// Analog Comparator
	ACSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	DIDR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x7f)))
	ACSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))

	// I/O Port
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x23)))
	PORTC = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	DDRC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	PINC  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	PORTD = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	DDRD  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	PIND  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	PORTE = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	DDRE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))
	PINE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))

	// Timer/Counter, 8-bit
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x47)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x46)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6e)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))

	// External Interrupts
	EICRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x69)))
	EIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))
	PCICR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x68)))
	PCMSK2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6d)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6c)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6b)))
	PCIFR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))

	// Serial Peripheral Interface
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x60)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))

	// CPU Registers
	PRR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x64)))
	OSCCAL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x66)))
	CLKPR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))
	SREG   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	SPMCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))
	MCUCR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x55)))
	MCUSR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	SMCR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	GPIOR2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	GPIOR1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	GPIOR0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BOOTSZ0 = 0x2 // Select boot size
	EXTENDED_BOOTSZ1 = 0x4 // Select boot size
	EXTENDED_BOOTRST = 0x1 // Boot Reset vector Enabled

	// HIGH
	HIGH_RSTDISBL  = 0x80 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_DWEN      = 0x40 // Debug Wire enable
	HIGH_SPIEN     = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON     = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE    = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BODLEVEL0 = 0x1  // Brown-out Detector trigger level
	HIGH_BODLEVEL1 = 0x2  // Brown-out Detector trigger level
	HIGH_BODLEVEL2 = 0x4  // Brown-out Detector trigger level

	// LOW
	LOW_CKDIV8     = 0x80 // Divide clock by 8 internally
	LOW_CKOUT      = 0x40 // Clock output on PORTB0
	LOW_SUT_CKSEL0 = 0x1  // Select Clock Source
	LOW_SUT_CKSEL1 = 0x2  // Select Clock Source
	LOW_SUT_CKSEL2 = 0x4  // Select Clock Source
	LOW_SUT_CKSEL3 = 0x8  // Select Clock Source
	LOW_SUT_CKSEL4 = 0x10 // Select Clock Source
	LOW_SUT_CKSEL5 = 0x20 // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0   = 0x1  // Memory Lock
	LOCKBIT_LB1   = 0x2  // Memory Lock
	LOCKBIT_BLB00 = 0x4  // Boot Loader Protection Mode
	LOCKBIT_BLB01 = 0x8  // Boot Loader Protection Mode
	LOCKBIT_BLB10 = 0x10 // Boot Loader Protection Mode
	LOCKBIT_BLB11 = 0x20 // Boot Loader Protection Mode
)

// Bitfields for USART: USART
const (
	// UCSR0A: USART Control and Status Register A
	UCSR0A_RXC0  = 0x80 // USART Receive Complete
	UCSR0A_TXC0  = 0x40 // USART Transmitt Complete
	UCSR0A_UDRE0 = 0x20 // USART Data Register Empty
	UCSR0A_FE0   = 0x10 // Framing Error
	UCSR0A_DOR0  = 0x8  // Data overRun
	UCSR0A_UPE0  = 0x4  // Parity Error
	UCSR0A_U2X0  = 0x2  // Double the USART transmission speed
	UCSR0A_MPCM0 = 0x1  // Multi-processor Communication Mode

	// UCSR0B: USART Control and Status Register B
	UCSR0B_RXCIE0 = 0x80 // RX Complete Interrupt Enable
	UCSR0B_TXCIE0 = 0x40 // TX Complete Interrupt Enable
	UCSR0B_UDRIE0 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR0B_RXEN0  = 0x10 // Receiver Enable
	UCSR0B_TXEN0  = 0x8  // Transmitter Enable
	UCSR0B_UCSZ02 = 0x4  // Character Size
	UCSR0B_RXB80  = 0x2  // Receive Data Bit 8
	UCSR0B_TXB80  = 0x1  // Transmit Data Bit 8

	// UCSR0C: USART Control and Status Register C
	UCSR0C_UMSEL00 = 0x40 // USART Mode Select
	UCSR0C_UMSEL01 = 0x80 // USART Mode Select
	UCSR0C_UPM00   = 0x10 // Parity Mode Bits
	UCSR0C_UPM01   = 0x20 // Parity Mode Bits
	UCSR0C_USBS0   = 0x8  // Stop Bit Select
	UCSR0C_UCSZ00  = 0x2  // Character Size
	UCSR0C_UCSZ01  = 0x4  // Character Size
	UCSR0C_UCPOL0  = 0x1  // Clock Polarity

	// UCSR0D: USART Control and Status Register D
	UCSR0D_RXSIE = 0x80 // RX Start Interrupt Enable
	UCSR0D_RXS   = 0x40 // RX Start
	UCSR0D_SFDE  = 0x20 // Start Frame Detection Enable
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWAMR: TWI (Slave) Address Mask Register
	TWAMR_TWAM0 = 0x2
	TWAMR_TWAM1 = 0x4
	TWAMR_TWAM2 = 0x8
	TWAMR_TWAM3 = 0x10
	TWAMR_TWAM4 = 0x20
	TWAMR_TWAM5 = 0x40
	TWAMR_TWAM6 = 0x80

	// TWCR: TWI Control Register
	TWCR_TWINT = 0x80 // TWI Interrupt Flag
	TWCR_TWEA  = 0x40 // TWI Enable Acknowledge Bit
	TWCR_TWSTA = 0x20 // TWI Start Condition Bit
	TWCR_TWSTO = 0x10 // TWI Stop Condition Bit
	TWCR_TWWC  = 0x8  // TWI Write Collition Flag
	TWCR_TWEN  = 0x4  // TWI Enable Bit
	TWCR_TWIE  = 0x1  // TWI Interrupt Enable

	// TWSR: TWI Status Register
	TWSR_TWS0  = 0x8  // TWI Status
	TWSR_TWS1  = 0x10 // TWI Status
	TWSR_TWS2  = 0x20 // TWI Status
	TWSR_TWS3  = 0x40 // TWI Status
	TWSR_TWS4  = 0x80 // TWI Status
	TWSR_TWPS0 = 0x1  // TWI Prescaler
	TWSR_TWPS1 = 0x2  // TWI Prescaler

	// TWAR: TWI (Slave) Address register
	TWAR_TWA0  = 0x2  // TWI (Slave) Address register Bits
	TWAR_TWA1  = 0x4  // TWI (Slave) Address register Bits
	TWAR_TWA2  = 0x8  // TWI (Slave) Address register Bits
	TWAR_TWA3  = 0x10 // TWI (Slave) Address register Bits
	TWAR_TWA4  = 0x20 // TWI (Slave) Address register Bits
	TWAR_TWA5  = 0x40 // TWI (Slave) Address register Bits
	TWAR_TWA6  = 0x80 // TWI (Slave) Address register Bits
	TWAR_TWGCE = 0x1  // TWI General Call Recognition Enable Bit
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output CompareB Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output CompareA Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0 = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1 = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0 = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1 = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_WGM10  = 0x1  // Waveform Generation Mode
	TCCR1A_WGM11  = 0x2  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10 = 0x8  // Waveform Generation Mode
	TCCR1B_WGM11 = 0x10 // Waveform Generation Mode
	TCCR1B_CS10  = 0x1  // Prescaler source of Timer/Counter 1
	TCCR1B_CS11  = 0x2  // Prescaler source of Timer/Counter 1
	TCCR1B_CS12  = 0x4  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A = 0x80
	TCCR1C_FOC1B = 0x40
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TIMSK2: Timer/Counter Interrupt Mask register
	TIMSK2_OCIE2B = 0x4 // Timer/Counter2 Output Compare Match B Interrupt Enable
	TIMSK2_OCIE2A = 0x2 // Timer/Counter2 Output Compare Match A Interrupt Enable
	TIMSK2_TOIE2  = 0x1 // Timer/Counter2 Overflow Interrupt Enable

	// TIFR2: Timer/Counter Interrupt Flag Register
	TIFR2_OCF2B = 0x4 // Output Compare Flag 2B
	TIFR2_OCF2A = 0x2 // Output Compare Flag 2A
	TIFR2_TOV2  = 0x1 // Timer/Counter2 Overflow Flag

	// TCCR2A: Timer/Counter2 Control Register A
	TCCR2A_COM2A0 = 0x40 // Compare Output Mode bits
	TCCR2A_COM2A1 = 0x80 // Compare Output Mode bits
	TCCR2A_COM2B0 = 0x10 // Compare Output Mode bits
	TCCR2A_COM2B1 = 0x20 // Compare Output Mode bits
	TCCR2A_WGM20  = 0x1  // Waveform Genration Mode
	TCCR2A_WGM21  = 0x2  // Waveform Genration Mode

	// TCCR2B: Timer/Counter2 Control Register B
	TCCR2B_FOC2A = 0x80 // Force Output Compare A
	TCCR2B_FOC2B = 0x40 // Force Output Compare B
	TCCR2B_WGM22 = 0x8  // Waveform Generation Mode
	TCCR2B_CS20  = 0x1  // Clock Select bits
	TCCR2B_CS21  = 0x2  // Clock Select bits
	TCCR2B_CS22  = 0x4  // Clock Select bits

	// ASSR: Asynchronous Status Register
	ASSR_EXCLK   = 0x40 // Enable External Clock Input
	ASSR_AS2     = 0x20 // Asynchronous Timer/Counter2
	ASSR_TCN2UB  = 0x10 // Timer/Counter2 Update Busy
	ASSR_OCR2AUB = 0x8  // Output Compare Register2 Update Busy
	ASSR_OCR2BUB = 0x4  // Output Compare Register 2 Update Busy
	ASSR_TCR2AUB = 0x2  // Timer/Counter Control Register2 Update Busy
	ASSR_TCR2BUB = 0x1  // Timer/Counter Control Register2 Update Busy
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS0 = 0x40 // Reference Selection Bits
	ADMUX_REFS1 = 0x80 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX0  = 0x1  // Analog Channel Selection Bits
	ADMUX_MUX1  = 0x2  // Analog Channel Selection Bits
	ADMUX_MUX2  = 0x4  // Analog Channel Selection Bits
	ADMUX_MUX3  = 0x8  // Analog Channel Selection Bits

	// ADCSRA: The ADC Control and Status register A
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC  Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC  Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC  Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC  Prescaler Select Bits

	// ADCSRB: The ADC Control and Status register B
	ADCSRB_ACME  = 0x40
	ADCSRB_ADTS0 = 0x1 // ADC Auto Trigger Source bits
	ADCSRB_ADTS1 = 0x2 // ADC Auto Trigger Source bits
	ADCSRB_ADTS2 = 0x4 // ADC Auto Trigger Source bits

	// DIDR0: Digital Input Disable Register
	DIDR0_ADC5D = 0x20
	DIDR0_ADC4D = 0x10
	DIDR0_ADC3D = 0x8
	DIDR0_ADC2D = 0x4
	DIDR0_ADC1D = 0x2
	DIDR0_ADC0D = 0x1
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD   = 0x80 // Analog Comparator Disable
	ACSR_ACBG  = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO   = 0x20 // Analog Compare Output
	ACSR_ACI   = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE  = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC  = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS0 = 0x1  // Analog Comparator Interrupt Mode Select bits
	ACSR_ACIS1 = 0x2  // Analog Comparator Interrupt Mode Select bits

	// DIDR1: Digital Input Disable Register 1
	DIDR1_AIN1D = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1 // AIN0 Digital Input Disable

	// ACSRB: Analog Comparator Status Register B
	ACSRB_ACOE = 0x1 // Analog Comparator Output Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8
	TCCR0B_CS00  = 0x1 // Clock Select
	TCCR0B_CS01  = 0x2 // Clock Select
	TCCR0B_CS02  = 0x4 // Clock Select

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A0 = 0x40 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0A1 = 0x80 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0B0 = 0x10 // Compare Output Mode, Fast PWm
	TCCR0A_COM0B1 = 0x20 // Compare Output Mode, Fast PWm
	TCCR0A_WGM00  = 0x1  // Waveform Generation Mode
	TCCR0A_WGM01  = 0x2  // Waveform Generation Mode

	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag 0B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag 0A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register
	EICRA_ISC10 = 0x4 // External Interrupt Sense Control 1 Bits
	EICRA_ISC11 = 0x8 // External Interrupt Sense Control 1 Bits
	EICRA_ISC00 = 0x1 // External Interrupt Sense Control 0 Bits
	EICRA_ISC01 = 0x2 // External Interrupt Sense Control 0 Bits

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT0 = 0x1 // External Interrupt Request 1 Enable
	EIMSK_INT1 = 0x2 // External Interrupt Request 1 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF0 = 0x1 // External Interrupt Flags
	EIFR_INTF1 = 0x2 // External Interrupt Flags

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE0 = 0x1 // Pin Change Interrupt Enables
	PCICR_PCIE1 = 0x2 // Pin Change Interrupt Enables
	PCICR_PCIE2 = 0x4 // Pin Change Interrupt Enables

	// PCMSK2: Pin Change Mask Register 2
	PCMSK2_PCINT0 = 0x1  // Pin Change Enable Masks
	PCMSK2_PCINT1 = 0x2  // Pin Change Enable Masks
	PCMSK2_PCINT2 = 0x4  // Pin Change Enable Masks
	PCMSK2_PCINT3 = 0x8  // Pin Change Enable Masks
	PCMSK2_PCINT4 = 0x10 // Pin Change Enable Masks
	PCMSK2_PCINT5 = 0x20 // Pin Change Enable Masks
	PCMSK2_PCINT6 = 0x40 // Pin Change Enable Masks
	PCMSK2_PCINT7 = 0x80 // Pin Change Enable Masks

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT0 = 0x1  // Pin Change Enable Masks
	PCMSK1_PCINT1 = 0x2  // Pin Change Enable Masks
	PCMSK1_PCINT2 = 0x4  // Pin Change Enable Masks
	PCMSK1_PCINT3 = 0x8  // Pin Change Enable Masks
	PCMSK1_PCINT4 = 0x10 // Pin Change Enable Masks
	PCMSK1_PCINT5 = 0x20 // Pin Change Enable Masks
	PCMSK1_PCINT6 = 0x40 // Pin Change Enable Masks

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT0 = 0x1  // Pin Change Enable Masks
	PCMSK0_PCINT1 = 0x2  // Pin Change Enable Masks
	PCMSK0_PCINT2 = 0x4  // Pin Change Enable Masks
	PCMSK0_PCINT3 = 0x8  // Pin Change Enable Masks
	PCMSK0_PCINT4 = 0x10 // Pin Change Enable Masks
	PCMSK0_PCINT5 = 0x20 // Pin Change Enable Masks
	PCMSK0_PCINT6 = 0x40 // Pin Change Enable Masks
	PCMSK0_PCINT7 = 0x80 // Pin Change Enable Masks

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF0 = 0x1 // Pin Change Interrupt Flags
	PCIFR_PCIF1 = 0x2 // Pin Change Interrupt Flags
	PCIFR_PCIF2 = 0x4 // Pin Change Interrupt Flags
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit

	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR0 = 0x1  // SPI Clock Rate Selects
	SPCR_SPR1 = 0x2  // SPI Clock Rate Selects
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP0 = 0x1  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP1 = 0x2  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP2 = 0x4  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP3 = 0x20 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EEPM0 = 0x10 // EEPROM Programming Mode Bits
	EECR_EEPM1 = 0x20 // EEPROM Programming Mode Bits
	EECR_EERIE = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMPE = 0x4  // EEPROM Master Write Enable
	EECR_EEPE  = 0x2  // EEPROM Write Enable
	EECR_EERE  = 0x1  // EEPROM Read Enable
)

// Bitfields for CPU: CPU Registers
const (
	// PRR: Power Reduction Register
	PRR_PRTWI    = 0x80 // Power Reduction TWI
	PRR_PRTIM2   = 0x40 // Power Reduction Timer/Counter2
	PRR_PRTIM0   = 0x20 // Power Reduction Timer/Counter0
	PRR_PRTIM1   = 0x8  // Power Reduction Timer/Counter1
	PRR_PRSPI    = 0x4  // Power Reduction Serial Peripheral Interface
	PRR_PRUSART0 = 0x2  // Power Reduction USART
	PRR_PRADC    = 0x1  // Power Reduction ADC

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL0 = 0x1  // Oscillator Calibration
	OSCCAL_OSCCAL1 = 0x2  // Oscillator Calibration
	OSCCAL_OSCCAL2 = 0x4  // Oscillator Calibration
	OSCCAL_OSCCAL3 = 0x8  // Oscillator Calibration
	OSCCAL_OSCCAL4 = 0x10 // Oscillator Calibration
	OSCCAL_OSCCAL5 = 0x20 // Oscillator Calibration
	OSCCAL_OSCCAL6 = 0x40 // Oscillator Calibration
	OSCCAL_OSCCAL7 = 0x80 // Oscillator Calibration

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS0 = 0x1  // Clock Prescaler Select Bits
	CLKPR_CLKPS1 = 0x2  // Clock Prescaler Select Bits
	CLKPR_CLKPS2 = 0x4  // Clock Prescaler Select Bits
	CLKPR_CLKPS3 = 0x8  // Clock Prescaler Select Bits

	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// SPMCSR: Store Program Memory Control and Status Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read-While-Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read-While-Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory

	// MCUCR: MCU Control Register
	MCUCR_BODS  = 0x40 // BOD Sleep
	MCUCR_BODSE = 0x20 // BOD Sleep Enable
	MCUCR_PUD   = 0x10
	MCUCR_IVSEL = 0x2
	MCUCR_IVCE  = 0x1

	// MCUSR: MCU Status Register
	MCUSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2 // External Reset Flag
	MCUSR_PORF  = 0x1 // Power-on reset flag

	// SMCR: Sleep Mode Control Register
	SMCR_SM0 = 0x2 // Sleep Mode Select Bits
	SMCR_SM1 = 0x4 // Sleep Mode Select Bits
	SMCR_SM2 = 0x8 // Sleep Mode Select Bits
	SMCR_SE  = 0x1 // Sleep Enable
)
