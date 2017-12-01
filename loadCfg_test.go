package v8runner

import (
	"path"

	"./v8dumpMode"

	"./v8constants"
	"./v8tools"
	. "gopkg.in/check.v1"
)

var _ = Suite(&тестыНаЗагрузкуКонфигурацииИзФайла{})
var _ = Suite(&тестыНаЗагрузкуКонфигурацииИзФайлов{})

type тестыНаЗагрузкуКонфигурацииИзФайла struct {
	conf                            *конфигуратор
	КаталогЗагрузки                 string
	ПутьКФайлуКофигурации           string
	ПутьКФайлуСпискаОбъектов        string
	ФорматВыгрузки                  string
	ОбновитьФайлИнформацииОВыгрузке bool
}

type тестыНаЗагрузкуКонфигурацииИзФайлов struct {
	conf                            *конфигуратор
	КаталогЗагрузки                 string
	ПутьКФайлуКофигурации           string
	ПутьКФайлуСпискаОбъектов        string
	ФорматВыгрузки                  string
	ОбновитьФайлИнформацииОВыгрузке bool
}

func (s *тестыНаЗагрузкуКонфигурацииИзФайла) SetUpSuite(c *C) {

	s.ПутьКФайлуКофигурации = path.Join(pwd, "tests", "fixtures", "ТестовыйФайлКонфигурации.cf")

}

func (s *тестыНаЗагрузкуКонфигурацииИзФайлов) SetUpSuite(c *C) {

	s.ПутьКФайлуКофигурации = path.Join(pwd, "tests", "fixtures", "ТестовыйФайлКонфигурации.cf")
	s.КаталогЗагрузки = v8tools.ВременныйКаталог()
	s.ФорматВыгрузки = РежимВыгрузкиКонфигурации.Иерархический

	conf := НовыйКонфигуратор()
	errLoad := conf.ЗагрузитьКонфигурациюИзФайла(s.ПутьКФайлуКофигурации)
	c.Assert(errLoad, IsNil, Commentf("Не удалось выполнить загрузку конфигурации: %s", s.ПутьКФайлуКофигурации))

	err := conf.ВыгрузитьКонфигурациюСРежимомВыгрузки(s.КаталогЗагрузки, s.ФорматВыгрузки)
	c.Assert(err, IsNil, Commentf("Не удалось выполностьб выгрузку конфигурации в каталог: %s", s.КаталогЗагрузки))

	xmlFile := path.Join(s.КаталогЗагрузки, v8constants.СonfiguratuonXml)
	_, err = v8tools.Exists(xmlFile)

	c.Assert(err, IsNil, Commentf("Файл с выгруженной конфигурацией не найден: %s", xmlFile))

}

func (s *тестыНаЗагрузкуКонфигурацииИзФайла) SetUpTest(c *C) {

	s.conf = НовыйКонфигуратор()
	s.ОбновитьФайлИнформацииОВыгрузке = false

}

func (s *тестыНаЗагрузкуКонфигурацииИзФайлов) SetUpTest(c *C) {

	s.conf = НовыйКонфигуратор()

}

func (s *тестыНаЗагрузкуКонфигурацииИзФайла) TearDownSuite(c *C) {
	v8tools.ОчиститьВременныйКаталог()
}

func (s *тестыНаЗагрузкуКонфигурацииИзФайлов) TearDownSuite(c *C) {
	v8tools.ОчиститьВременныйКаталог()
}

func (s *тестыНаЗагрузкуКонфигурацииИзФайлов) TestКонфигуратор_loadConfigFromFiles(c *C) {

	err := s.conf.loadConfigFromFiles(s.КаталогЗагрузки, s.ПутьКФайлуСпискаОбъектов, s.ФорматВыгрузки, s.ОбновитьФайлИнформацииОВыгрузке)
	c.Assert(err, IsNil)

}

func (s *тестыНаЗагрузкуКонфигурацииИзФайла) TestКонфигуратор_loadCfg(c *C) {

	err := s.conf.loadCfg(s.ПутьКФайлуКофигурации)
	c.Assert(err, IsNil)
}

func (s *тестыНаЗагрузкуКонфигурацииИзФайла) TestКонфигуратор_ЗагрузитьКонфигурациюИзФайла(c *C) {

	err := s.conf.ЗагрузитьКонфигурациюИзФайла(s.ПутьКФайлуКофигурации)
	c.Assert(err, IsNil)
}

func (s *тестыНаЗагрузкуКонфигурацииИзФайлов) TestКонфигуратор_ЗагрузитьКонфигурациюИзФайлов(c *C) {

	err := s.conf.ЗагрузитьКонфигурациюИзФайлов(s.КаталогЗагрузки)
	c.Assert(err, IsNil)
}
