export class CommonDataSource<T> implements IDataSource {
  private listeners: DataChangeListener[] = [];
  protected items: T[] = []

  public totalCount(): number {
    return this.items.length
  }

  public getData(index: number): T | void {
    return this.items[index]
  };

  public getList(): T[] {
    return this.items
  }

  public clearData() {
    this.items = [];
    this.notifyDataReload()
  }

  public resetData(data: T | T[]) {
    this.items = [];
    if (Array.isArray(data)) {
      this.items.push(...data);
    } else {
      this.items.push(data);
    }
    this.notifyDataReload()
  }

  public setDataArray(data: T[]) {
    this.items = data
    this.notifyDataReload()
  }

  // 向列表追加数据
  public pushData(data: T | T[]): void {
    let fromIndex = this.items.length
    if (Array.isArray(data)) {
      this.items.push(...data)
    } else {
      this.items.push(data)
    }
    this.notifyDatasetAdd(fromIndex, this.items.length - fromIndex)
  }


  public changeItem(index: number) {
    this.notifyDataChange(index)
  }

  public deleteData(index: number): void {
    this.items.splice(index, 1);
    this.notifyDataDelete(index);
  }

  registerDataChangeListener(listener: DataChangeListener): void {
    if (this.listeners.indexOf(listener) < 0) {
      this.listeners.push(listener);
    }
  }

  unregisterDataChangeListener(listener: DataChangeListener): void {
    const pos = this.listeners.indexOf(listener);
    if (pos >= 0) {
      this.listeners.splice(pos, 1);
    }
  }

  notifyDataReload(): void {
    this.listeners.forEach((listener: DataChangeListener) => {
      listener.onDatasetChange([{ type: DataOperationType.RELOAD }]);
    });
  }

  notifyDatasetAdd(insertIndex: number, addCount: number): void {
    this.listeners.forEach(listener => {
      listener.onDatasetChange([{ type: DataOperationType.ADD, index: insertIndex, count: addCount }]);
    })
  }

  notifyDataChange(changeIndex: number): void {
    this.listeners.forEach((listener: DataChangeListener) => {
      listener.onDatasetChange([{ type: DataOperationType.CHANGE, index: changeIndex }]);
    });
  }

  notifyDataDelete(deleteIndex: number): void {
    this.listeners.forEach((listener: DataChangeListener) => {
      listener.onDatasetChange([{ type: DataOperationType.DELETE, index: deleteIndex }]);
    });
  }
}