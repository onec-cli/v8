package v8

import (
	"github.com/v8platform/designer"
)

func DumpIB(file string) designer.DumpIBOptions {

	command := designer.DumpIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

func RestoreIB(file string) designer.RestoreIBOptions {

	command := designer.RestoreIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

func IBRestoreIntegrity() designer.IBRestoreIntegrityOptions {

	return designer.IBRestoreIntegrityOptions{
		Designer: designer.NewDesigner(),
	}
}

func UpdateDBCfg(server bool, Dynamic bool) designer.UpdateDBCfgOptions {

	command := designer.UpdateDBCfgOptions{
		Designer: designer.NewDesigner(),
		Server:   server,
		Dynamic:  Dynamic,
	}

	return command

}

func UpdateDBExtensionCfg(extension string, server bool, Dynamic bool) designer.UpdateDBCfgOptions {

	command := UpdateDBCfg(server, Dynamic)
	command.Extension = extension

	return command

}

func CreateFileInfobase(file string) designer.CreateFileInfoBaseOptions {

	command := designer.CreateFileInfoBaseOptions{
		File: file,
	}

	return command

}
