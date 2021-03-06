// Automatically generated file. DO NOT EDIT.
// Generated by gen-device-avr.go from ATtiny841.atdf, see http://packs.download.atmel.com/

// +build avr,attiny841

// Device information for the ATtiny841.
package avr

import (
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

// Some information about this device.
const (
	DEVICE = "ATtiny841"
	ARCH   = "AVR8"
	FAMILY = "tinyAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset and Watchdog Reset
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_PCINT0       = 2  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 3  // Pin Change Interrupt Request 1
	IRQ_WDT          = 4  // Watchdog Time-out Interrupt
	IRQ_TIMER1_CAPT  = 5  // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 6  // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 7  // Timer/Counter1 Compare Match B
	IRQ_TIMER1_OVF   = 8  // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 9  // TimerCounter0 Compare Match A
	IRQ_TIMER0_COMPB = 10 // TimerCounter0 Compare Match B
	IRQ_TIMER0_OVF   = 11 // Timer/Couner0 Overflow
	IRQ_ANA_COMP0    = 12 // Analog Comparator 0
	IRQ_ADC          = 13 // ADC Conversion Complete
	IRQ_EE_RDY       = 14 // EEPROM Ready
	IRQ_ANA_COMP1    = 15 // Analog Comparator 1
	IRQ_TIMER2_CAPT  = 16 // Timer/Counter2 Capture Event
	IRQ_TIMER2_COMPA = 17 // Timer/Counter2 Compare Match A
	IRQ_TIMER2_COMPB = 18 // Timer/Counter2 Compare Match B
	IRQ_TIMER2_OVF   = 19 // Timer/Counter2 Overflow
	IRQ_SPI          = 20 // Serial Peripheral Interface
	IRQ_USART0_START = 21 // USART0, Start
	IRQ_USART0_RX    = 22 // USART0, Rx Complete
	IRQ_USART0_UDRE  = 23 // USART0 Data Register Empty
	IRQ_USART0_TX    = 24 // USART0, Tx Complete
	IRQ_USART1_START = 25 // USART1, Start
	IRQ_USART1_RX    = 26 // USART1, Rx Complete
	IRQ_USART1_UDRE  = 27 // USART1 Data Register Empty
	IRQ_USART1_TX    = 28 // USART1, Tx Complete
	IRQ_TWI_SLAVE    = 29 // Two-wire Serial Interface
	IRQ_max          = 29 // Highest interrupt number on this device.
)

// Map interrupt numbers to function names.
// These aren't real calls, they're removed by the compiler.
var (
	_ = interrupt.Register(IRQ_RESET, "__vector_RESET")
	_ = interrupt.Register(IRQ_INT0, "__vector_INT0")
	_ = interrupt.Register(IRQ_PCINT0, "__vector_PCINT0")
	_ = interrupt.Register(IRQ_PCINT1, "__vector_PCINT1")
	_ = interrupt.Register(IRQ_WDT, "__vector_WDT")
	_ = interrupt.Register(IRQ_TIMER1_CAPT, "__vector_TIMER1_CAPT")
	_ = interrupt.Register(IRQ_TIMER1_COMPA, "__vector_TIMER1_COMPA")
	_ = interrupt.Register(IRQ_TIMER1_COMPB, "__vector_TIMER1_COMPB")
	_ = interrupt.Register(IRQ_TIMER1_OVF, "__vector_TIMER1_OVF")
	_ = interrupt.Register(IRQ_TIMER0_COMPA, "__vector_TIMER0_COMPA")
	_ = interrupt.Register(IRQ_TIMER0_COMPB, "__vector_TIMER0_COMPB")
	_ = interrupt.Register(IRQ_TIMER0_OVF, "__vector_TIMER0_OVF")
	_ = interrupt.Register(IRQ_ANA_COMP0, "__vector_ANA_COMP0")
	_ = interrupt.Register(IRQ_ADC, "__vector_ADC")
	_ = interrupt.Register(IRQ_EE_RDY, "__vector_EE_RDY")
	_ = interrupt.Register(IRQ_ANA_COMP1, "__vector_ANA_COMP1")
	_ = interrupt.Register(IRQ_TIMER2_CAPT, "__vector_TIMER2_CAPT")
	_ = interrupt.Register(IRQ_TIMER2_COMPA, "__vector_TIMER2_COMPA")
	_ = interrupt.Register(IRQ_TIMER2_COMPB, "__vector_TIMER2_COMPB")
	_ = interrupt.Register(IRQ_TIMER2_OVF, "__vector_TIMER2_OVF")
	_ = interrupt.Register(IRQ_SPI, "__vector_SPI")
	_ = interrupt.Register(IRQ_USART0_START, "__vector_USART0_START")
	_ = interrupt.Register(IRQ_USART0_RX, "__vector_USART0_RX")
	_ = interrupt.Register(IRQ_USART0_UDRE, "__vector_USART0_UDRE")
	_ = interrupt.Register(IRQ_USART0_TX, "__vector_USART0_TX")
	_ = interrupt.Register(IRQ_USART1_START, "__vector_USART1_START")
	_ = interrupt.Register(IRQ_USART1_RX, "__vector_USART1_RX")
	_ = interrupt.Register(IRQ_USART1_UDRE, "__vector_USART1_UDRE")
	_ = interrupt.Register(IRQ_USART1_TX, "__vector_USART1_TX")
	_ = interrupt.Register(IRQ_TWI_SLAVE, "__vector_TWI_SLAVE")
)

// Peripherals.
var (
	// Fuses
	EXTENDED = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2)))
	HIGH     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x1)))
	LOW      = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// Lockbits
	LOCKBIT = (*volatile.Register8)(unsafe.Pointer(uintptr(0x0)))

	// I/O Port
	PUEB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x62)))
	DDRB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x37)))
	PINB  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x36)))
	PORTB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x38)))
	PUEA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x63)))
	PORTA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3b)))
	DDRA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3a)))
	PINA  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x39)))
	PHDE  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x6a)))

	// USART
	UDR1   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x90)))
	UCSR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x96)))
	UCSR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x95)))
	UCSR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x94)))
	UCSR1D = (*volatile.Register8)(unsafe.Pointer(uintptr(0x93)))
	UBRR1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x91)))
	UBRR1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x92)))
	UDR0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x80)))
	UCSR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x86)))
	UCSR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x85)))
	UCSR0C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x84)))
	UCSR0D = (*volatile.Register8)(unsafe.Pointer(uintptr(0x83)))
	UBRR0L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x81)))
	UBRR0H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x82)))

	// Watchdog Timer
	WDTCSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0x41)))

	// Two Wire Serial Interface
	TWSCRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa5)))
	TWSCRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa4)))
	TWSSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa3)))
	TWSA   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa2)))
	TWSD   = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa0)))
	TWSAM  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xa1)))

	// Analog-to-Digital Converter
	ADMUXA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x29)))
	ADMUXB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x28)))
	ADCSRA = (*volatile.Register8)(unsafe.Pointer(uintptr(0x25)))
	ADCL   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x26)))
	ADCH   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x27)))
	ADCSRB = (*volatile.Register8)(unsafe.Pointer(uintptr(0x24)))
	DIDR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x61)))
	DIDR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x60)))

	// Analog Comparator
	ACSR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2b)))
	ACSR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2a)))
	ACSR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2d)))
	ACSR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2c)))

	// EEPROM
	EEARL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3e)))
	EEARH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3f)))
	EEDR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3d)))
	EECR  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x3c)))

	// Timer/Counter, 16-bit
	TIMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2f)))
	TIFR1  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x2e)))
	TCCR1A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4f)))
	TCCR1B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4e)))
	TCCR1C = (*volatile.Register8)(unsafe.Pointer(uintptr(0x42)))
	TCNT1L = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4c)))
	TCNT1H = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4d)))
	OCR1AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4a)))
	OCR1AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x4b)))
	OCR1BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0x48)))
	OCR1BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0x49)))
	ICR1L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x44)))
	ICR1H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x45)))
	TIMSK2 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x31)))
	TIFR2  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x30)))
	TCCR2A = (*volatile.Register8)(unsafe.Pointer(uintptr(0xca)))
	TCCR2B = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc9)))
	TCCR2C = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc8)))
	TCNT2L = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc6)))
	TCNT2H = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc7)))
	OCR2AL = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc4)))
	OCR2AH = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc5)))
	OCR2BL = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc2)))
	OCR2BH = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc3)))
	ICR2L  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc0)))
	ICR2H  = (*volatile.Register8)(unsafe.Pointer(uintptr(0xc1)))

	// Timer/Counter, 8-bit
	TIMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x59)))
	TIFR0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x58)))
	TCCR0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x50)))
	TCCR0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x53)))
	TCNT0  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x52)))
	OCR0A  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x56)))
	OCR0B  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5c)))

	// External Interrupts
	GIMSK  = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5b)))
	GIFR   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5a)))
	PCMSK1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x40)))
	PCMSK0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x32)))

	// CPU Registers
	PRR       = (*volatile.Register8)(unsafe.Pointer(uintptr(0x70)))
	CCP       = (*volatile.Register8)(unsafe.Pointer(uintptr(0x71)))
	CLKPR     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x73)))
	CLKCR     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x72)))
	SREG      = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5f)))
	SPL       = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5d)))
	SPH       = (*volatile.Register8)(unsafe.Pointer(uintptr(0x5e)))
	MCUSR     = (*volatile.Register8)(unsafe.Pointer(uintptr(0x54)))
	GPIOR2    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x35)))
	GPIOR1    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x34)))
	GPIOR0    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x33)))
	SPMCSR    = (*volatile.Register8)(unsafe.Pointer(uintptr(0x57)))
	OSCCAL0   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x74)))
	OSCCAL1   = (*volatile.Register8)(unsafe.Pointer(uintptr(0x77)))
	OSCTCAL0A = (*volatile.Register8)(unsafe.Pointer(uintptr(0x75)))
	OSCTCAL0B = (*volatile.Register8)(unsafe.Pointer(uintptr(0x76)))

	// Timer/Counter Output Compare Pin
	TOCPMSA1 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x68)))
	TOCPMSA0 = (*volatile.Register8)(unsafe.Pointer(uintptr(0x67)))
	TOCPMCOE = (*volatile.Register8)(unsafe.Pointer(uintptr(0x66)))

	// Serial Peripheral Interface
	SPCR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb2)))
	SPSR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb1)))
	SPDR = (*volatile.Register8)(unsafe.Pointer(uintptr(0xb0)))
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_ULPOSCSEL0 = 0x20 // Frequency selection for internal ULP oscillator. The selection only affects system clock, watchdog and reset timeout always use 32 kHz clock.
	EXTENDED_ULPOSCSEL1 = 0x40 // Frequency selection for internal ULP oscillator. The selection only affects system clock, watchdog and reset timeout always use 32 kHz clock.
	EXTENDED_ULPOSCSEL2 = 0x80 // Frequency selection for internal ULP oscillator. The selection only affects system clock, watchdog and reset timeout always use 32 kHz clock.
	EXTENDED_BODPD0     = 0x8  // BOD mode of operation when the device is in sleep mode
	EXTENDED_BODPD1     = 0x10 // BOD mode of operation when the device is in sleep mode
	EXTENDED_BODACT0    = 0x2  // BOD mode of operation when the device is active or idle
	EXTENDED_BODACT1    = 0x4  // BOD mode of operation when the device is active or idle
	EXTENDED_SELFPRGEN  = 0x1  // Self Programming enable

	// HIGH
	HIGH_RSTDISBL  = 0x80 // Reset Disabled (Enable PC2 as i/o pin)
	HIGH_DWEN      = 0x40 // Debug Wire enable
	HIGH_SPIEN     = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON     = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE    = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BODLEVEL0 = 0x1  // Brown-out Detector trigger level
	HIGH_BODLEVEL1 = 0x2  // Brown-out Detector trigger level
	HIGH_BODLEVEL2 = 0x4  // Brown-out Detector trigger level

	// LOW
	LOW_CKDIV8     = 0x80 // Divide clock by 8 internally
	LOW_CKOUT      = 0x40 // Clock output on PORTC2
	LOW_SUT_CKSEL0 = 0x1  // Select Clock Source
	LOW_SUT_CKSEL1 = 0x2  // Select Clock Source
	LOW_SUT_CKSEL2 = 0x4  // Select Clock Source
	LOW_SUT_CKSEL3 = 0x8  // Select Clock Source
	LOW_SUT_CKSEL4 = 0x10 // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB0 = 0x1 // Memory Lock
	LOCKBIT_LB1 = 0x2 // Memory Lock
)

// Bitfields for PORT: I/O Port
const (
	// PHDE: Port High Drive Enable Register
	PHDE_PHDEA0 = 0x1 // PortA High Drive Enable
	PHDE_PHDEA1 = 0x2 // PortA High Drive Enable
)

// Bitfields for USART: USART
const (
	// UCSR1A: USART Control and Status Register A
	UCSR1A_RXC1  = 0x80 // USART Receive Complete
	UCSR1A_TXC1  = 0x40 // USART Transmitt Complete
	UCSR1A_UDRE1 = 0x20 // USART Data Register Empty
	UCSR1A_FE1   = 0x10 // Framing Error
	UCSR1A_DOR1  = 0x8  // Data overRun
	UCSR1A_UPE1  = 0x4  // Parity Error
	UCSR1A_U2X1  = 0x2  // Double the USART transmission speed
	UCSR1A_MPCM1 = 0x1  // Multi-processor Communication Mode

	// UCSR1B: USART Control and Status Register B
	UCSR1B_RXCIE1 = 0x80 // RX Complete Interrupt Enable
	UCSR1B_TXCIE1 = 0x40 // TX Complete Interrupt Enable
	UCSR1B_UDRIE1 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR1B_RXEN1  = 0x10 // Receiver Enable
	UCSR1B_TXEN1  = 0x8  // Transmitter Enable
	UCSR1B_UCSZ12 = 0x4  // Character Size
	UCSR1B_RXB81  = 0x2  // Receive Data Bit 8
	UCSR1B_TXB81  = 0x1  // Transmit Data Bit 8

	// UCSR1C: USART Control and Status Register C
	UCSR1C_UMSEL10 = 0x40 // USART Mode Select
	UCSR1C_UMSEL11 = 0x80 // USART Mode Select
	UCSR1C_UPM10   = 0x10 // Parity Mode Bits
	UCSR1C_UPM11   = 0x20 // Parity Mode Bits
	UCSR1C_USBS1   = 0x8  // Stop Bit Select
	UCSR1C_UCSZ10  = 0x2  // Character Size
	UCSR1C_UCSZ11  = 0x4  // Character Size
	UCSR1C_UCPOL1  = 0x1  // Clock Polarity

	// UCSR1D: USART Control and Status Register D
	UCSR1D_RXSIE1 = 0x80 // USART RX Start Interrupt Enable
	UCSR1D_RXS1   = 0x40 // USART RX Start Flag
	UCSR1D_SFDE1  = 0x20 // USART RX Start Frame Detection Enable

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
	UCSR0D_RXSIE0 = 0x80 // USART RX Start Interrupt Enable
	UCSR0D_RXS0   = 0x40 // USART RX Start Flag
	UCSR0D_SFDE0  = 0x20 // USART RX Start Frame Detection Enable
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control and Status Register
	WDTCSR_WDIF = 0x80 // Watchdog Timer Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timer Interrupt Enable
	WDTCSR_WDP0 = 0x1  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP1 = 0x2  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP2 = 0x4  // Watchdog Timer Prescaler Bits
	WDTCSR_WDP3 = 0x20 // Watchdog Timer Prescaler Bits
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWSCRA: TWI Slave Control Register A
	TWSCRA_TWSHE  = 0x80 // TWI SDA Hold Time Enable
	TWSCRA_TWDIE  = 0x20 // TWI Data Interrupt Enable
	TWSCRA_TWASIE = 0x10 // TWI Address/Stop Interrupt Enable
	TWSCRA_TWEN   = 0x8  // Two-Wire Interface Enable
	TWSCRA_TWSIE  = 0x4  // TWI Stop Interrupt Enable
	TWSCRA_TWPME  = 0x2  // TWI Promiscuous Mode Enable
	TWSCRA_TWSME  = 0x1  // TWI Smart Mode Enable

	// TWSCRB: TWI Slave Control Register B
	TWSCRB_TWHNM  = 0x8 // TWI High Noise Mode
	TWSCRB_TWAA   = 0x4 // TWI Acknowledge Action
	TWSCRB_TWCMD0 = 0x1
	TWSCRB_TWCMD1 = 0x2

	// TWSSRA: TWI Slave Status Register A
	TWSSRA_TWDIF  = 0x80 // TWI Data Interrupt Flag.
	TWSSRA_TWASIF = 0x40 // TWI Address/Stop Interrupt Flag
	TWSSRA_TWCH   = 0x20 // TWI Clock Hold
	TWSSRA_TWRA   = 0x10 // TWI Receive Acknowledge
	TWSSRA_TWC    = 0x8  // TWI Collision
	TWSSRA_TWBE   = 0x4  // TWI Bus Error
	TWSSRA_TWDIR  = 0x2  // TWI Read/Write Direction
	TWSSRA_TWAS   = 0x1  // TWI Address or Stop

	// TWSD: TWI Slave Data Register
	TWSD_TWSD0 = 0x1  // TWI slave data bit
	TWSD_TWSD1 = 0x2  // TWI slave data bit
	TWSD_TWSD2 = 0x4  // TWI slave data bit
	TWSD_TWSD3 = 0x8  // TWI slave data bit
	TWSD_TWSD4 = 0x10 // TWI slave data bit
	TWSD_TWSD5 = 0x20 // TWI slave data bit
	TWSD_TWSD6 = 0x40 // TWI slave data bit
	TWSD_TWSD7 = 0x80 // TWI slave data bit

	// TWSAM: TWI Slave Address Mask Register
	TWSAM_TWSAM0 = 0x2  // TWI Address Mask Bits
	TWSAM_TWSAM1 = 0x4  // TWI Address Mask Bits
	TWSAM_TWSAM2 = 0x8  // TWI Address Mask Bits
	TWSAM_TWSAM3 = 0x10 // TWI Address Mask Bits
	TWSAM_TWSAM4 = 0x20 // TWI Address Mask Bits
	TWSAM_TWSAM5 = 0x40 // TWI Address Mask Bits
	TWSAM_TWSAM6 = 0x80 // TWI Address Mask Bits
	TWSAM_TWAE   = 0x1  // TWI Address Enable
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUXA: The ADC multiplexer Selection Register A
	ADMUXA_MUX0 = 0x1  // Analog Channel and Gain Selection Bits
	ADMUXA_MUX1 = 0x2  // Analog Channel and Gain Selection Bits
	ADMUXA_MUX2 = 0x4  // Analog Channel and Gain Selection Bits
	ADMUXA_MUX3 = 0x8  // Analog Channel and Gain Selection Bits
	ADMUXA_MUX4 = 0x10 // Analog Channel and Gain Selection Bits
	ADMUXA_MUX5 = 0x20 // Analog Channel and Gain Selection Bits

	// ADMUXB: The ADC multiplexer Selection Register B
	ADMUXB_REFS0 = 0x20 // Reference Selection Bits
	ADMUXB_REFS1 = 0x40 // Reference Selection Bits
	ADMUXB_REFS2 = 0x80 // Reference Selection Bits
	ADMUXB_GSEL0 = 0x1  // Gain Selection Bits
	ADMUXB_GSEL1 = 0x2  // Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS0 = 0x1  // ADC Prescaler Select Bits
	ADCSRA_ADPS1 = 0x2  // ADC Prescaler Select Bits
	ADCSRA_ADPS2 = 0x4  // ADC Prescaler Select Bits

	// ADCSRB: ADC Control and Status Register B
	ADCSRB_ADLAR = 0x8
	ADCSRB_ADTS0 = 0x1 // ADC Auto Trigger Sources
	ADCSRB_ADTS1 = 0x2 // ADC Auto Trigger Sources
	ADCSRB_ADTS2 = 0x4 // ADC Auto Trigger Sources

	// DIDR1: Digital Input Disable Register 1
	DIDR1_ADC9D  = 0x8 // ADC9 Digital Input Disable
	DIDR1_ADC8D  = 0x4 // ADC8 Digital input Disable
	DIDR1_ADC10D = 0x2 // ADC10 Digital input Disable
	DIDR1_ADC11D = 0x1 // ADC11 Digital input Disable

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC7D = 0x80 // ADC7 Digital input Disable
	DIDR0_ADC6D = 0x40 // ADC6 Digital input Disable
	DIDR0_ADC5D = 0x20 // ADC5 Digital input Disable
	DIDR0_ADC4D = 0x10 // ADC4/AIN11 Digital input Disable
	DIDR0_ADC3D = 0x8  // ADC3/AIN10 Digital Input Disable
	DIDR0_ADC2D = 0x4  // ADC2/AIN01 Digital input Disable
	DIDR0_ADC1D = 0x2  // ADC1/AIN00 Digital input Disable
	DIDR0_ADC0D = 0x1  // ADC0/AREF Digital input Disable
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR0B: Analog Comparator 0 Control And Status Register B
	ACSR0B_HSEL0   = 0x80 // Analog Comparator 0 Hysteresis Select
	ACSR0B_HLEV0   = 0x40 // Analog Comparator 0 Hysteresis Level
	ACSR0B_ACOE0   = 0x10 // Analog Comparator 0 Output Pin Enable
	ACSR0B_ACNMUX0 = 0x4  // Analog Comparator 0 Negative Input Multiplexer
	ACSR0B_ACNMUX1 = 0x8  // Analog Comparator 0 Negative Input Multiplexer
	ACSR0B_ACPMUX0 = 0x1  // Analog Comparator 0 Positive Input Multiplexer Bits 1:0
	ACSR0B_ACPMUX1 = 0x2  // Analog Comparator 0 Positive Input Multiplexer Bits 1:0

	// ACSR0A: Analog Comparator 0 Control And Status Register A
	ACSR0A_ACD0    = 0x80 // Analog Comparator 0 Disable
	ACSR0A_ACPMUX2 = 0x40 // Analog Comparator 0 Positive Input Multiplexer Bit 2
	ACSR0A_ACO0    = 0x20 // Analog Comparator 0 Output
	ACSR0A_ACI0    = 0x10 // Analog Comparator 0 Interrupt Flag
	ACSR0A_ACIE0   = 0x8  // Analog Comparator 0 Interrupt Enable
	ACSR0A_ACIC0   = 0x4  // Analog Comparator 0 Input Capture Enable
	ACSR0A_ACIS00  = 0x1  // Analog Comparator 0 Interrupt Mode Select bits
	ACSR0A_ACIS01  = 0x2  // Analog Comparator 0 Interrupt Mode Select bits

	// ACSR1B: Analog Comparator 1 Control And Status Register B
	ACSR1B_HSEL1 = 0x80 // Analog Comparator 1 Hysteresis Select
	ACSR1B_HLEV1 = 0x40 // Analog Comparator 1 Hysteresis Level
	ACSR1B_ACOE1 = 0x10 // Analog Comparator 1 Output Pin Enable
	ACSR1B_ACME1 = 0x4  // Analog Comparator 1 Multiplexer Enable

	// ACSR1A: Analog Comparator 1 Control And Status Register A
	ACSR1A_ACD1   = 0x80 // Analog Comparator 1 Disable
	ACSR1A_ACBG1  = 0x40 // Analog Comparator 1 Bandgap Select
	ACSR1A_ACO1   = 0x20 // Analog Comparator 1 Output
	ACSR1A_ACI1   = 0x10 // Analog Comparator 1 Interrupt Flag
	ACSR1A_ACIE1  = 0x8  // Analog Comparator 1 Interrupt Enable
	ACSR1A_ACIC1  = 0x4  // Analog Comparator 1 Input Capture Enable
	ACSR1A_ACIS10 = 0x1  // Analog Comparator 1 Interrupt Mode Select bits
	ACSR1A_ACIS11 = 0x2  // Analog Comparator 1 Interrupt Mode Select bits
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

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TIMSK1: Timer/Counter1 Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Timer/Counter1 Input Capture Flag
	TIFR1_OCF1B = 0x4  // Timer/Counter1 Output Compare B Match Flag
	TIFR1_OCF1A = 0x2  // Timer/Counter1 Output Compare A Match Flag
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A0 = 0x40 // Compare Output Mode 1A, bits
	TCCR1A_COM1A1 = 0x80 // Compare Output Mode 1A, bits
	TCCR1A_COM1B0 = 0x10 // Compare Output Mode 1B, bits
	TCCR1A_COM1B1 = 0x20 // Compare Output Mode 1B, bits
	TCCR1A_WGM10  = 0x1  // Pulse Width Modulator Select Bits
	TCCR1A_WGM11  = 0x2  // Pulse Width Modulator Select Bits

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM10 = 0x8  // Waveform Generation Mode Bits
	TCCR1B_WGM11 = 0x10 // Waveform Generation Mode Bits
	TCCR1B_CS10  = 0x1  // Clock Select bits
	TCCR1B_CS11  = 0x2  // Clock Select bits
	TCCR1B_CS12  = 0x4  // Clock Select bits

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare for Channel A
	TCCR1C_FOC1B = 0x40 // Force Output Compare for Channel B

	// TIMSK2: Timer/Counter2 Interrupt Mask Register
	TIMSK2_ICIE2  = 0x20 // Timer/Counter2 Input Capture Interrupt Enable
	TIMSK2_OCIE2B = 0x4  // Timer/Counter2 Output Compare B Match Interrupt Enable
	TIMSK2_OCIE2A = 0x2  // Timer/Counter2 Output Compare A Match Interrupt Enable
	TIMSK2_TOIE2  = 0x1  // Timer/Counter2 Overflow Interrupt Enable

	// TIFR2: Timer/Counter Interrupt Flag register
	TIFR2_ICF2  = 0x20 // Timer/Counter2 Input Capture Flag
	TIFR2_OCF2B = 0x4  // Timer/Counter2 Output Compare B Match Flag
	TIFR2_OCF2A = 0x2  // Timer/Counter2 Output Compare A Match Flag
	TIFR2_TOV2  = 0x1  // Timer/Counter2 Overflow Flag

	// TCCR2A: Timer/Counter2 Control Register A
	TCCR2A_COM2A0 = 0x40 // Compare Output Mode 2A, bits
	TCCR2A_COM2A1 = 0x80 // Compare Output Mode 2A, bits
	TCCR2A_COM2B0 = 0x10 // Compare Output Mode 2B, bits
	TCCR2A_COM2B1 = 0x20 // Compare Output Mode 2B, bits
	TCCR2A_WGM20  = 0x1  // Pulse Width Modulator Select Bits
	TCCR2A_WGM21  = 0x2  // Pulse Width Modulator Select Bits

	// TCCR2B: Timer/Counter2 Control Register B
	TCCR2B_ICNC2 = 0x80 // Input Capture 2 Noise Canceler
	TCCR2B_ICES2 = 0x40 // Input Capture 2 Edge Select
	TCCR2B_WGM20 = 0x8  // Waveform Generation Mode Bits
	TCCR2B_WGM21 = 0x10 // Waveform Generation Mode Bits
	TCCR2B_CS20  = 0x1  // Clock Select bits
	TCCR2B_CS21  = 0x2  // Clock Select bits
	TCCR2B_CS22  = 0x4  // Clock Select bits

	// TCCR2C: Timer/Counter2 Control Register C
	TCCR2C_FOC2A = 0x80 // Force Output Compare for Channel A
	TCCR2C_FOC2B = 0x40 // Force Output Compare for Channel B
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TIMSK0: Timer/Counter Interrupt Mask Register
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag Register
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A0 = 0x40 // Compare Match Output A Mode bits
	TCCR0A_COM0A1 = 0x80 // Compare Match Output A Mode bits
	TCCR0A_COM0B0 = 0x10 // Compare Match Output B Mode bits
	TCCR0A_COM0B1 = 0x20 // Compare Match Output B Mode bits
	TCCR0A_WGM00  = 0x1  // Waveform Generation Mode bits
	TCCR0A_WGM01  = 0x2  // Waveform Generation Mode bits

	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8  // Waveform Generation Mode bit 2
	TCCR0B_CS00  = 0x1  // Clock Select bits
	TCCR0B_CS01  = 0x2  // Clock Select bits
	TCCR0B_CS02  = 0x4  // Clock Select bits
)

// Bitfields for EXINT: External Interrupts
const (
	// GIMSK: General Interrupt Mask Register
	GIMSK_INT0  = 0x40 // External Interrupt Request 0 Enable
	GIMSK_PCIE0 = 0x10 // Pin Change Interrupt Enables
	GIMSK_PCIE1 = 0x20 // Pin Change Interrupt Enables

	// GIFR: General Interrupt Flag register
	GIFR_INTF0 = 0x40 // External Interrupt Flag 0
	GIFR_PCIF0 = 0x10 // Pin Change Interrupt Flags
	GIFR_PCIF1 = 0x20 // Pin Change Interrupt Flags

	// PCMSK1: Pin Change Enable Mask 1
	PCMSK1_PCINT8  = 0x1 // Pin Change Enable Mask 1 Bit 0
	PCMSK1_PCINT9  = 0x2 // Pin Change Enable Mask 1 Bit 1
	PCMSK1_PCINT10 = 0x4 // Pin Change Enable Mask 1 Bit 2
	PCMSK1_PCINT11 = 0x8 // Pin Change Enable Mask 1 Bit 3

	// PCMSK0: Pin Change Enable Mask 0
	PCMSK0_PCINT0 = 0x1  // Pin Change Enable Mask 0 Bit 0
	PCMSK0_PCINT1 = 0x2  // Pin Change Enable Mask 0 Bit 1
	PCMSK0_PCINT2 = 0x4  // Pin Change Enable Mask 0 Bit 2
	PCMSK0_PCINT3 = 0x8  // Pin Change Enable Mask 0 Bit 3
	PCMSK0_PCINT4 = 0x10 // Pin Change Enable Mask 0 Bit 4
	PCMSK0_PCINT5 = 0x20 // Pin Change Enable Mask 0 Bit 5
	PCMSK0_PCINT6 = 0x40 // Pin Change Enable Mask 0 Bit 6
	PCMSK0_PCINT7 = 0x80 // Pin Change Enable Mask 0 Bit 7
)

// Bitfields for CPU: CPU Registers
const (
	// PRR: Power Reduction Register
	PRR_PRTWI    = 0x80 // Power Reduction TWI
	PRR_PRUSART1 = 0x40 // Power Reduction USART1
	PRR_PRUSART0 = 0x20 // Power Reduction USART0
	PRR_PRSPI    = 0x10 // Power Reduction SPI
	PRR_PRTIM2   = 0x8  // Power Reduction Timer/Counter2
	PRR_PRTIM1   = 0x4  // Power Reduction Timer/Counter1
	PRR_PRTIM0   = 0x2  // Power Reduction Timer/Counter0
	PRR_PRADC    = 0x1  // Power Reduction ADC

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPS0 = 0x1 // Clock Prescaler Select Bits
	CLKPR_CLKPS1 = 0x2 // Clock Prescaler Select Bits
	CLKPR_CLKPS2 = 0x4 // Clock Prescaler Select Bits
	CLKPR_CLKPS3 = 0x8 // Clock Prescaler Select Bits

	// CLKCR: Clock Control Register
	CLKCR_OSCRDY = 0x80 // Oscillator Ready
	CLKCR_CSTR   = 0x40 // Clock Switch Trigger
	CLKCR_CKOUTC = 0x20 // Clock Output (Copy). Active low.
	CLKCR_SUT    = 0x10 // Start-up Time
	CLKCR_CKSEL0 = 0x1  // Clock Select Bits
	CLKCR_CKSEL1 = 0x2  // Clock Select Bits
	CLKCR_CKSEL2 = 0x4  // Clock Select Bits
	CLKCR_CKSEL3 = 0x8  // Clock Select Bits

	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// MCUSR: MCU Status Register
	MCUSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2 // External Reset Flag
	MCUSR_PORF  = 0x1 // Power-on reset flag

	// SPMCSR: Store Program Memory Control and Status Register
	SPMCSR_RSIG  = 0x20 // Read Device Signature Imprint Table
	SPMCSR_CTPB  = 0x10 // Clear Temporary Page Buffer
	SPMCSR_RFLB  = 0x8  // Read Fuse and Lock Bits
	SPMCSR_PGWRT = 0x4  // Page Write
	SPMCSR_PGERS = 0x2  // Page Erase
	SPMCSR_SPMEN = 0x1  // Store program Memory Enable
)

// Bitfields for TOCPM: Timer/Counter Output Compare Pin
const (
	// TOCPMSA1: Timer Output Compare Pin Mux Selection 1
	TOCPMSA1_TOCC7S0 = 0x40 // Timer Output Compare Channel 7 Selection Bits
	TOCPMSA1_TOCC7S1 = 0x80 // Timer Output Compare Channel 7 Selection Bits
	TOCPMSA1_TOCC6S0 = 0x10 // Timer Output Compare Channel 6 Selection Bits
	TOCPMSA1_TOCC6S1 = 0x20 // Timer Output Compare Channel 6 Selection Bits
	TOCPMSA1_TOCC5S0 = 0x4  // Timer Output Compare Channel 5 Selection Bits
	TOCPMSA1_TOCC5S1 = 0x8  // Timer Output Compare Channel 5 Selection Bits
	TOCPMSA1_TOCC4S0 = 0x1  // Timer Output Compare Channel 4 Selection Bits
	TOCPMSA1_TOCC4S1 = 0x2  // Timer Output Compare Channel 4 Selection Bits

	// TOCPMSA0: Timer Output Compare Pin Mux Selection 0
	TOCPMSA0_TOCC3S0 = 0x40 // Timer Output Compare Channel 3 Selection Bits
	TOCPMSA0_TOCC3S1 = 0x80 // Timer Output Compare Channel 3 Selection Bits
	TOCPMSA0_TOCC2S0 = 0x10 // Timer Output Compare Channel 2 Selection Bits
	TOCPMSA0_TOCC2S1 = 0x20 // Timer Output Compare Channel 2 Selection Bits
	TOCPMSA0_TOCC1S0 = 0x4  // Timer Output Compare Channel 1 Selection Bits
	TOCPMSA0_TOCC1S1 = 0x8  // Timer Output Compare Channel 1 Selection Bits
	TOCPMSA0_TOCC0S0 = 0x1  // Timer Output Compare Channel 0 Selection Bits
	TOCPMSA0_TOCC0S1 = 0x2  // Timer Output Compare Channel 0 Selection Bits

	// TOCPMCOE: Timer Output Compare Pin Mux Channel Output Enable
	TOCPMCOE_TOCC7OE = 0x80 // Timer Output Compare Channel 7 Output Enable
	TOCPMCOE_TOCC6OE = 0x40 // Timer Output Compare Channel 6 Output Enable
	TOCPMCOE_TOCC5OE = 0x20 // Timer Output Compare Channel 5 Output Enable
	TOCPMCOE_TOCC4OE = 0x10 // Timer Output Compare Channel 4 Output Enable
	TOCPMCOE_TOCC3OE = 0x8  // Timer Output Compare Channel 3 Output Enable
	TOCPMCOE_TOCC2OE = 0x4  // Timer Output Compare Channel 2 Output Enable
	TOCPMCOE_TOCC1OE = 0x2  // Timer Output Compare Channel 1 Output Enable
	TOCPMCOE_TOCC0OE = 0x1  // Timer Output Compare Channel 0 Output Enable
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR0 = 0x1  // SPI Clock Rate Selects
	SPCR_SPR1 = 0x2  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit
)
