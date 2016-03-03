package zfs

// GetZpool retrieves a single ZFS zpool by name.
func GetZpool(name string) (*Zpool, error) {
	out, err := zpool("get", "-p", "all", name)
	if err != nil {
		return nil, err
	}

	// there is no -H
	out = out[1:]

	z := &Zpool{Name: name}
	for _, line := range out {
		if err := z.parseLine(line); err != nil {
			return nil, err
		}
	}

	return z, nil
}
