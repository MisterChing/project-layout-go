package bootstrap

type BeforeServerFunc func() error
type AfterServerFunc func()

type Bootloader struct {
    BeforeServerFuncs []BeforeServerFunc
    AfterServerFuncs  []AfterServerFunc
}

func NewBootStrap() *Bootloader {
    return new(Bootloader)
}

func (boot *Bootloader) AddBeforeServerFunc(fns ...BeforeServerFunc) {
    for _, fn := range fns {
        boot.BeforeServerFuncs = append(boot.BeforeServerFuncs, fn)
    }
}

func (boot *Bootloader) AddAfterServerFunc(fns ...AfterServerFunc) {
    for _, fn := range fns {
        boot.AfterServerFuncs = append(boot.AfterServerFuncs, fn)
    }
}

func (boot *Bootloader) SetUp() error {
    for _, fn := range boot.BeforeServerFuncs {
        if err := fn(); err != nil {
            return err
        }
    }
    return nil
}

func (boot *Bootloader) Destroy() {
    for _, fn := range boot.AfterServerFuncs {
        fn()
    }
}
